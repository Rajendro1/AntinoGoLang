package mysqldb

import "github.com/Rajendro1/AntinoGoLang/config"

var DbAndTableCreation = `
CREATE DATABASE ` + config.DB_NAME + `;
USE ` + config.DB_NAME + `;

CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    content TEXT,
    author VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

`
