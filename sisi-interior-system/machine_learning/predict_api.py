from flask import Flask, request, jsonify
import pickle
import pandas as pd

app = Flask(__name__)

# load model
data = pickle.load(open("model_estimasi.pkl", "rb"))
model = data["model"]
columns = data["columns"]

@app.route("/predict", methods=["POST"])
def predict():
    req = request.json

    df = pd.DataFrame([req])

    # encoding (HARUS SAMA seperti training)
    df = pd.get_dummies(df)
    df = df.reindex(columns=columns, fill_value=0)

    hasil = model.predict(df)[0]

    return jsonify({
        "estimasi": float(hasil)
    })

app.run(port=5000)