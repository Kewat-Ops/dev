# Read secret at runtime
with open("/run/secrets/db_password") as f:
    db_pass = f.read().strip()
    print("DB password:", db_pass)

from flask import Flask
from prometheus_flask_exporter import PrometheusMetrics

# OpenTelemetry imports
from opentelemetry import trace
from opentelemetry.instrumentation.flask import FlaskInstrumentor
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.exporter.jaeger.thrift import JaegerExporter

# Setup tracing
trace.set_tracer_provider(TracerProvider())
jaeger_exporter = JaegerExporter(agent_host_name="jaeger", agent_port=6831)
trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(jaeger_exporter))

app = Flask(__name__)
FlaskInstrumentor().instrument_app(app)

# Prometheus metrics
metrics = PrometheusMetrics(app)

@app.route("/python")
def hello():
    # Create a span for tracing
    tracer = trace.get_tracer(__name__)
    with tracer.start_as_current_span("hello_route"):
        return f"Hello from Python! Secret is: {db_pass}"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)

