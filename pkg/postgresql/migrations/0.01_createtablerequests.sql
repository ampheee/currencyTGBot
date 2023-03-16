CREATE TABLE requests (
    request_id INT PRIMARY KEY,
    request_type VARCHAR(15),
    request_time TIMESTAMP,
    request_response TEXT
);

DROP TABLE requests;