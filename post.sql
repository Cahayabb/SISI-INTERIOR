CREATE TABLE admin (
id SERIAL PRIMARY KEY,
username VARCHAR(50) UNIQUE,
password VARCHAR(255)
);

INSERT INTO admin (username, password)
VALUES ('admin', '123456');

SELECT * FROM admin;

CREATE TABLE portfolio (
    id SERIAL PRIMARY KEY,
    judul_proyek VARCHAR(150) NOT NULL,
    deskripsi TEXT,
    gambar TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);