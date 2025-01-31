from flask import Flask, jsonify

app = Flask(__name__)

@app.route("/")
def home():
    return jsonify({"email": "lawsonredeye@mail.com",
                    "current_datetime": "Eze your datetime go dey here",
                    "github_url": "https://github.com/lawsonredeye/my_repo"
                }), 200