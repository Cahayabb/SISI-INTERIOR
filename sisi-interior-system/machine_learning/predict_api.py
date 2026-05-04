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
    try:
        req = request.json

        required_fields = [
                "luas_area",
                "tingkat_kerumitan",
                "durasi_pengerjaan",
                "jenis_ruangan",
                "jenis_pekerjaan",
                "spesifikasi_design"
            ]

        for field in required_fields:
                if field not in req:
                    return jsonify({"error": f"{field} wajib diisi"}), 400

        df = pd.DataFrame([req])

        # encoding (HARUS SAMA seperti training)
        df = pd.get_dummies(df)
        df = df.reindex(columns=columns, fill_value=0)

        hasil = model.predict(df)[0]

        return jsonify({
            "estimasi": float(hasil)
        })
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == "__main__":
    app.run(port=5000, debug=True)