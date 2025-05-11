from flask import Flask, jsonify
from multiprocessing import Value

# future work -- test atomicity with Redis support. This would probably be in a
# sepearate container, if anything.
counter = Value('i', 0)

app = Flask(__name__)


@app.route('/')
def index():
    with counter.get_lock():
        counter.value += 1
        out = counter.value

    return jsonify(count=out)


if __name__ == "__main__":
    app.run(processes=1, host='0.0.0.0', port=8000)
