CREATE TABLE IF NOT EXISTS users (
    u_id INT PRIMARY KEY NOT NULL,
    username VARCHAR(64) NOT NULL,
    firstname VARCHAR(32),
    lastname VARCHAR(32)
);

CREATE TABLE IF NOT EXISTS requests (
    u_id INT,
    FOREIGN KEY (u_id) REFERENCES users(u_id),
    r_id SERIAL PRIMARY KEY,
    r_type VARCHAR(16),
    r_args VARCHAR(32),
    r_time TIMESTAMP,
);
