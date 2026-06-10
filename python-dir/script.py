# code to read secrets at runtime inside container
with open("/run/secrets/db_password") as f:
    db_pass = f.read().strip()
    print("DB password:", db_pass)


# main script of apps with routing
from flask import Flask

# using prometheus exporter (like blackbox)
from prometheus_flask_exporter import PrometheusMetrics

app = Flask(__name__)
metrics = PrometheusMetrics(app)

@app.route("/python")
def hello():
    return "Hello from Python!"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
