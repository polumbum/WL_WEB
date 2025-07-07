SELECT CURRENT_DATE

SELECT * FROM competitions c 


CREATE OR REPLACE FUNCTION get_upcoming_competitions()
RETURNS TABLE (
    id UUID,
    name TEXT,
    city TEXT,
    address TEXT,
    beg_date DATE,
    end_date DATE,
    age TEXT,
    min_sports_category TEXT,
    antidoping BOOL
) AS $$
BEGIN
    RETURN QUERY
    SELECT *
    FROM competitions
    WHERE competitions.end_date >= CURRENT_DATE;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM get_upcoming_competitions();

CREATE OR REPLACE FUNCTION get_upcoming_tcamps()
RETURNS TABLE (
    id UUID,
    city TEXT,
    address TEXT,
    beg_date DATE,
    end_date DATE
) AS $$
BEGIN
    RETURN QUERY
    SELECT *
    FROM t_camps
    WHERE t_camps.end_date >= CURRENT_DATE;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM get_upcoming_tcamps();

CREATE OR REPLACE FUNCTION get_upcoming_competitions_sm(sm_id UUID) 
RETURNS TABLE (
    id UUID,
	name TEXT,
	city TEXT,
	address TEXT,
	beg_date DATE,
	end_date DATE,
	age TEXT,
	min_sports_category TEXT,
	antidoping BOOL
) AS $$
BEGIN
    RETURN QUERY
    SELECT c.id, c.name, c.city, c.address, c.beg_date, c.end_date, c.age, c.min_sports_category, c.antidoping
    FROM competitions c
    JOIN comp_applications ca ON c.id = ca.competition_id
    WHERE ca.sportsman_id = sm_id
    AND c.beg_date > CURRENT_DATE;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM get_upcoming_competitions_sm('f8f26e6d-3e36-416a-9ff0-253876fc1115');

select * from sportsmen;



select * from sportsmen
delete from sportsmen s 

INSERT INTO sportsmen (surname, name, patronymic, birthday, moscow_team, sports_category, gender)
VALUES ('A', 'B', 'C', '10.12.2002', true, 'КМС', true);

drop table sportsmen 


INSERT INTO coaches (surname, name, patronymic, experience, birthday, gender)
VALUES ('A', 'A', 'A', 2, '10.12.2010', true);

select * from coaches
-- drop table coaches cascade

select sportsmen from sportsmen
inner join sportsmen_coaches on sportsmen.id = sportsmen_coaches.sportsman_id
inner join  coaches on coaches.id = sportsmen_coaches.coach_id and coaches.id = 'afc4caca-796e-471a-8bae-0b3bb857660a'

select * from sportsmen s 
select * from sportsmen_coaches sc 

INSERT INTO sportsmen_coaches (sportsman_id, coach_id)
VALUES ('519b2a4c-d872-4511-9410-325103b67143', 'afc4caca-796e-471a-8bae-0b3bb857660a')

INSERT INTO sportsmen_coaches (sportsman_id, coach_id)
VALUES ('125ca7de-b934-4b29-b5eb-6bc95ba3e71b', 'f8f26e6d-3e36-416a-9ff0-253876fc1115')
-- smID
-- 519b2a4c-d872-4511-9410-325103b67143

-- compID
-- 4012c4ac-fd43-4669-b7f4-4e4c1cabe527
-- 41798cdc-891a-48e3-8686-7d703f2ada1b
select * from users u 
select * from results 
CREATE TABLE results (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	sportsman_id UUID REFERENCES sportsmen(id),
	competition_id UUID REFERENCES competitions(id),
	weight_category INT NOT NULL CHECK (weight_category > 0),
	snatch INT NOT NULL,
	clean_and_jerk INT NOT NULL,
	place INT NOT NULL
);

INSERT INTO results (sportsman_id, competition_id, weight_category, snatch, clean_and_jerk, place)
VALUES ('519b2a4c-d872-4511-9410-325103b67143', '4012c4ac-fd43-4669-b7f4-4e4c1cabe527', 59, 60, 70, 1);
INSERT INTO results (sportsman_id, competition_id, weight_category, snatch, clean_and_jerk, place)
VALUES ('519b2a4c-d872-4511-9410-325103b67143', '41798cdc-891a-48e3-8686-7d703f2ada1b', 59, 60, 70, 1);


drop table sportsmen_coaches 

select * from comp_accesses ca 

INSERT INTO comp_accesses (sportsman_id, institution, validity)
VALUES ('519b2a4c-d872-4511-9410-325103b67143', 'INSTITUTION', '2025-10-23');

drop table comp_accesses 


select * from antidopings a 

drop table antidopings 

drop table t_camps cascade

select * from competitions
drop table competitions 

INSERT INTO competitions (name, city, address, beg_date, end_date, age, min_sports_category, antidoping)
VALUES ('A', 'B', 'C', '2023-12-10', '2023-12-12', 'мужчины, женщины', 'I', true);

INSERT INTO competitions (name, city, address, beg_date, end_date, age, min_sports_category, antidoping)
VALUES ('A', 'B', 'C', '2023-12-10', '2023-12-12', 'мужчины, женщины', 'II', false);

drop table competitions

select * from t_camps tc 

select * from t_camp_applications tca 

drop table t_camp_applications 

select * from comp_applications 
drop table comp_applications 

select * from results

drop table results 

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

drop table users

select * from users
select * from sportsmen where id LIKE '95606ce3-f4cd-469c-b4e1-6bb286aafabe'