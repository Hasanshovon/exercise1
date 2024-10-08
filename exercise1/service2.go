package main

import (
    "encoding/json"
    "net"
    "net/http"
    "os/exec"
    "strings"
)

type SystemInfo struct {
    IPAddress   string `json:"ip_address"`
    Processes   string `json:"processes"`
    DiskSpace   string `json:"disk_space"`
    Uptime      string `json:"uptime"`
}

func getSystemInfo() SystemInfo {
    // Get IP address
    addrs, _ := net.InterfaceAddrs()
    var ip string
    for _, addr := range addrs {
        if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
            if ipNet.IP.To4() != nil {
                ip = ipNet.IP.String()
                break
            }
        }
    }

    // Get running processes
    out, _ := exec.Command("ps", "-ax").Output()
    processes := string(out)

    // Get available disk space
    out, _ = exec.Command("df", "-h").Output()
    diskSpace := string(out)

    // Get time since last boot
    out, _ = exec.Command("uptime", "-p").Output()
    uptime := string(out)

    return SystemInfo{
        IPAddress: ip,
        Processes: processes,
        DiskSpace: diskSpace,
        Uptime:    strings.TrimSpace(uptime),
    }
}

func service2Handler(w http.ResponseWriter, r *http.Request) {
    info := getSystemInfo()
    json.NewEncoder(w).Encode(info)
}

func main() {
    http.HandleFunc("/service2", service2Handler)
    http.ListenAndServe(":5001", nil)
}
