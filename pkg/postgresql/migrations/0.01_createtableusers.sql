CREATE TABLE users (
    id INT PRIMARY KEY,
    total_requests INT,
    FOREIGN KEY (id) REFERENCES requests(id)
);
DROP TABLE users;