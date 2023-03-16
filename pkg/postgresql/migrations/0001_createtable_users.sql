CREATE TABLE requests (
    request_id INT PRIMARY KEY,
    request_type VARCHAR(15),
    request_time TIMESTAMP,
    request_response TEXT
);

CREATE TABLE users (
    id INT PRIMARY KEY,
    total_requests INT,
    FOREIGN KEY (id) REFERENCES requests (id)
);
---- create above / drop below ----
DROP TABLE users;
DROP TABLE requests;