import pandas as pd
import psycopg2
from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split
from sklearn.metrics import mean_absolute_error
import pickle

# koneksi database
conn = psycopg2.connect(
    host="localhost",
    database="sisiinterior",
    user="postgres",
    password="123",
    port="5433"
)

query = """
SELECT 
    luas_area, 
    tingkat_kerumitan, 
    durasi_pengerjaan,
    jenis_ruangan,
    jenis_pekerjaan,
    spesifikasi_design,
    harga_proyek
FROM data_proyek
"""
df = pd.read_sql(query, conn)

df = df.dropna()

# mapping string --> angka
df['tingkat_kerumitan'] = df['tingkat_kerumitan'].str.strip().map({
    'Rendah': 1,
    'Sedang': 2,
    'Tinggi': 3
})

# memastikan numeric
df['durasi_pengerjaan'] = pd.to_numeric(df['durasi_pengerjaan'], errors='coerce')

df = df.dropna()

# encoding
df = pd.get_dummies(df, columns=[
    'jenis_ruangan',
    'jenis_pekerjaan',
    'spesifikasi_design'
])

X = df.drop('harga_proyek', axis=1)
y = df['harga_proyek']

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2)

model = LinearRegression()
model.fit(X_train, y_train)

pred = model.predict(X_test)

print("MAE:", mean_absolute_error(y_test, pred))
# simpan
pickle.dump({
    "model": model,
    "columns": X.columns
}, open("model_estimasi.pkl","wb"))

print("Model siap digunakan")

