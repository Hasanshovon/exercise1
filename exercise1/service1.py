import socket
import subprocess
import requests
from flask import Flask, jsonify

app = Flask(__name__)

def get_system_info():
    # Get IP address
    hostname = socket.gethostname()
    ip_address = socket.gethostbyname(hostname)

    # Get running processes
    processes = subprocess.check_output(["ps", "-ax"]).decode("utf-8")

    # Get available disk space
    disk_space = subprocess.check_output(["df", "-h"]).decode("utf-8")

    # Get time since last boot
    uptime = subprocess.check_output(["uptime", "-p"]).decode("utf-8")

    return {
        "ip_address": ip_address,
        "processes": processes,
        "disk_space": disk_space,
        "uptime": uptime
    }

@app.route('/service1', methods=['GET'])
def service1_info():
    # Get information from Service2
    service2_info = requests.get('http://service2:5001/service2').json()
    
    # Get information from Service1
    service1_info = get_system_info()

    # Combine both
    response = {
        "Service1": service1_info,
        "Service2": service2_info
    }

    return jsonify(response)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8199)
