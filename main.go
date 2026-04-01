package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	port        = "8347"
	host        = "0.0.0.0"
	downloadDir = "/home/admin/snap-downloads"
)

func main() {
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		log.Fatalf("Failed to create download directory: %v", err)
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/download", handleDownload)

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Server starting on http://%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Snap Installer</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
        input[type="text"] { width: 100%; padding: 10px; margin: 10px 0; box-sizing: border-box; }
        button { padding: 10px 20px; background-color: #4CAF50; color: white; border: none; cursor: pointer; }
        button:hover { background-color: #45a049; }
        #status { margin-top: 20px; padding: 10px; border-radius: 5px; }
        .success { background-color: #d4edda; color: #155724; }
        .error { background-color: #f8d7da; color: #721c24; }
        .info { background-color: #d1ecf1; color: #0c5460; }
    </style>
</head>
<body>
    <h1>Snap Package Installer</h1>
    <form id="downloadForm">
        <label for="url">Enter Snap Package URL:</label>
        <input type="text" id="url" name="url" placeholder="https://example.com/package.snap" required>
        <button type="submit">Download and Install</button>
    </form>
    <div id="status"></div>

    <script>
        document.getElementById('downloadForm').addEventListener('submit', async (e) => {
            e.preventDefault();
            const url = document.getElementById('url').value;
            const statusDiv = document.getElementById('status');

            statusDiv.className = 'info';
            statusDiv.textContent = 'Downloading...';

            try {
                const response = await fetch('/download', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
                    body: 'url=' + encodeURIComponent(url)
                });

                const result = await response.json();

                if (response.ok) {
                    statusDiv.className = 'success';
                    statusDiv.textContent = result.message;
                } else {
                    statusDiv.className = 'error';
                    statusDiv.textContent = 'Error: ' + result.error;
                }
            } catch (error) {
                statusDiv.className = 'error';
                statusDiv.textContent = 'Error: ' + error.message;
            }
        });
    </script>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "URL is required"})
		return
	}

	//http://10.0.5.224:8111/repository/download/Snap_Cloudex/4369:id/cloudex_20260401_amd64.snap
	log.Printf("Downloading Snap Package URL: %s", url)

	filename := filepath.Base(url)
	filepath := filepath.Join(downloadDir, filename)

	if err := downloadFile(filepath, url); err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Download failed: %v", err)})
		return
	}

	defer os.Remove(filepath)
	if err := installSnap(filepath); err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Installation failed: %v", err)})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("Successfully downloaded and installed %s", filename)})
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url + "?guest")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func installSnap(filepath string) error {
	cmd := exec.Command("snap", "install", filepath, "--dangerous")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, output)
	}
	log.Printf("Snap installed successfully: %s", output)
	return nil
}

func respondJSON(w http.ResponseWriter, status int, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"message":"%s","error":"%s"}`, data["message"], data["error"])
}
