CREATE DATABASE IF NOT EXISTS galaveg DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER IF NOT EXISTS 'galaveg_user'@'localhost' IDENTIFIED BY 'galaveg_password';

GRANT ALL ON galaveg.* TO 'galaveg_user'@'localhost';