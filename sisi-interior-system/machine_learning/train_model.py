import pandas as pd
import psycopg2
from sklearn.linear_model import LinearRegression
import pickle

# koneksi database
conn = psycopg2.connect(
    host="localhost",
    database="sisiinterior",
    user="postgres",
    password="123",
    port="5433"
)

print("Koneksi berhasil")

query = """
SELECT luas_area, tingkat_kerumitan, durasi_pengerjaan, harga_proyek
FROM data_proyek
"""

df = pd.read_sql(query, conn)

X = df[['luas_area','tingkat_kerumitan','durasi_pengerjaan']]
y = df['harga_proyek']

model = LinearRegression()
model.fit(X,y)

# simpan model
pickle.dump(model, open("model_estimasi.pkl","wb"))

print("Model berhasil dibuat")