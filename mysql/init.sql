CREATE TABLE IF NOT EXISTS courses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO courses (name) VALUES
('Docker for Beginners'),
('AWS DevOps Essentials'),
('Kubernetes Hands-On'),
('GitHub Actions CI/CD');