create table faces (
    id BIGSERIAL PRIMARY KEY,
    "aws_face_id" varchar(255)
  );
alter table faces add column email VARCHAR(255);
select * from faces;
delete from faces;
create table face_activity (
    id BIGSERIAL PRIMARY KEY,
    "aws_face_id" varchar(255),
    the_time TIMESTAMP
);
INSERT into face_activity (aws_face_id, the_time) VALUES
	('aws1', '2020-09-01 09:00:00'), ('aws1', '2020-09-01 09:00:02'), ('aws1', '2020-09-01 09:00:03'), ('aws1', '2020-09-01 09:00:04'), ('aws1', '2020-09-01 09:00:05'),
	('aws1', '2020-09-01 10:00:00'), ('aws1', '2020-09-01 10:00:02'), ('aws1', '2020-09-01 10:00:03'), ('aws1', '2020-09-01 10:00:04'), ('aws1', '2020-09-01 10:00:05'),
	('aws1', '2020-09-01 10:30:00'), ('aws1', '2020-09-01 10:30:02'), ('aws1', '2020-09-01 10:30:03'), ('aws1', '2020-09-01 10:30:04'), ('aws1', '2020-09-01 10:30:05'),
	('aws1', '2020-09-01 09:00:00'), ('aws1', '2020-09-01 09:00:02'), ('aws1', '2020-09-01 09:00:03'), ('aws1', '2020-09-01 09:00:04'), ('aws1', '2020-09-01 09:00:05'),
	('aws1', '2020-09-01 10:00:00'), ('aws1', '2020-09-01 10:00:02'), ('aws1', '2020-09-01 10:00:03'), ('aws1', '2020-09-01 10:00:04'), ('aws1', '2020-09-01 10:00:05'),
	('aws1', '2020-09-01 10:30:00'), ('aws1', '2020-09-01 10:30:02'), ('aws1', '2020-09-01 10:30:03'), ('aws1', '2020-09-01 10:30:04'), ('aws1', '2020-09-01 10:30:05')
;
INSERT into face_activity (aws_face_id, the_time) VALUES
	('aws1', '2020-09-01 09:00:00'), ('aws1', '2020-10-01 09:00:02'), ('aws1', '2020-10-01 09:00:03'), ('aws1', '2020-10-01 09:00:04'), ('aws1', '2020-10-01 09:00:05'),
	('aws1', '2020-10-01 10:00:00'), ('aws1', '2020-10-01 10:00:02'), ('aws1', '2020-10-01 10:00:03'), ('aws1', '2020-10-01 10:00:04'), ('aws1', '2020-10-01 10:00:05'),
	('aws1', '2020-10-01 10:30:00'), ('aws1', '2020-10-01 10:30:02'), ('aws1', '2020-10-01 10:30:03'), ('aws1', '2020-10-01 10:30:04'), ('aws1', '2020-10-01 10:30:05'),
	('aws1', '2020-10-01 09:00:00'), ('aws1', '2020-10-01 09:00:02'), ('aws1', '2020-10-01 09:00:03'), ('aws1', '2020-10-01 09:00:04'), ('aws1', '2020-10-01 09:00:05'),
	('aws1', '2020-10-01 10:00:00'), ('aws1', '2020-10-01 10:00:02'), ('aws1', '2020-10-01 10:00:03'), ('aws1', '2020-10-01 10:00:04'), ('aws1', '2020-10-01 10:00:05'),
	('aws1', '2020-10-01 10:30:00'), ('aws1', '2020-10-01 10:30:02'), ('aws1', '2020-10-01 10:30:03'), ('aws1', '2020-10-01 10:30:04'), ('aws1', '2020-09-01 10:30:05')
;
