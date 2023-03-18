CREATE TABLE IF NOT EXISTS users (
    u_id INT PRIMARY KEY UNIQUE NOT NULL,
    username VARCHAR(64) NOT NULL,
    name VARCHAR(32),
    surname VARCHAR(32)
);

CREATE TABLE IF NOT EXISTS requests (
    u_id INT,
    FOREIGN KEY (u_id) REFERENCES users(u_id),
    r_id SERIAL PRIMARY KEY,
    r_type VARCHAR(15),
    r_time TIMESTAMP,
    r_response TEXT
);
