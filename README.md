
# A simple HTTP Server

I design the service with MVC structure,
This service will handle for upload file like images, and this service have 2 routes


| Method  | Route |Description|
| ------------- |-------------|-------------|
| GET      | /ping     | check the service is up/down|
| GET / POST      | /upload     | user can upload file|


service will running on port :8000, you can access page directly like this
http://localhost:8000/upload
http://localhost:8000/files/sample.jpeg


every file success uploaded will insert to table log_files.
So you need a setup database like Mysql, i will share the DDL for that

```
CREATE DATABASE uploader;
CREATE TABLE log_files (
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    contentType varchar(10),
    size int,
    userAgent text,
    urlFile text,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

Implement by using net/http and lib mysql
