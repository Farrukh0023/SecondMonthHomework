create table students (
    id serial primary key,
    first_name varchar(50),
    last_name varchar(50),
    date_of_birth date,
    math_exam_score float,
    history_exam_score float,
    geography_exam_score float
);

insert into students (first_name, last_name, date_of_birth, math_exam_score, history_exam_score, geography_exam_score)
values (
('Farrux', 'Ashirmatov', '2000-01-01', 4.5, 3.2, 4.8),
('Samandar', 'Ubaydullayev', '2001-02-03', 3.8, 4.5, 3.2),
('Nodir', 'Salimboyev', '1999-05-15', 2.5, 4.8, 3.5),
('Nodir', 'Odilov', '2002-08-20', 4.2, 3.7, 4.1),
('Aziz', 'Qayumov', '2000-11-10', 3.0, 2.8, 3.9);
);

select * from students where (math_exam_score + history_exam_score + geography_exam_score) / 3 > 3.5;

select * from students where math_exam_score >= 3 and math_exam_score <= 5;

select * from students where history_exam_score < 4 or history_exam_score = 4.5;

select * from students where history_exam_score < 4 or history_exam_score = 4.5;



