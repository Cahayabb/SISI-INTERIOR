from flask import Flask, request, jsonify
from flask_cors import CORS
import pickle
import pandas as pd

app = Flask(__name__)
CORS(app) 

# load model
data = pickle.load(open("model_estimasi.pkl", "rb"))
model = data["model"]
columns = data["columns"]

@app.route("/predict", methods=["POST"])
def predict():
    try:
        req = request.get_json()

        print("DATA MASUK:", req)

        if not req:
            return jsonify({"error": "Request kosong"}), 400


        required_fields = [
                "luas_area",
                "tingkat_kerumitan",
                "durasi_pengerjaan",
                "jenis_ruangan",
                "jenis_proyek",
                "spesifikasi_design"
            ]

        # VALIDASI FIELD
        for field in required_fields:
            if field not in req or req[field] in [None, ""]:
                return jsonify({"error": f"{field} wajib diisi"}), 400

        # convert ke dataframe
        df = pd.DataFrame([req])

        # encoding (HARUS sama seperti training)
        df = pd.get_dummies(df)

        # samakan kolom dengan model
        df = df.reindex(columns=columns, fill_value=0)

        # prediksi
        hasil = model.predict(df)[0]

        print("HASIL:", hasil)  # DEBUG

        return jsonify({
            "estimasi": float(hasil)
        })

    except Exception as e:
        print("ERROR:", str(e))  # DEBUG
        return jsonify({"error": str(e)}), 500


if __name__ == "__main__":
    app.run(host="127.0.0.1", port=5000, debug=True)