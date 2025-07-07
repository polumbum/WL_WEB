CREATE TABLE sportsmen (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	surname TEXT NOT NULL,
	name TEXT NOT NULL,
	patronymic TEXT,
	birthday DATE NOT NULL,
	moscow_team BOOL NOT NULL, 
	sports_category TEXT,
	gender BOOL NOT NULL
);

CREATE TABLE coaches (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	surname TEXT NOT NULL,
	name TEXT NOT NULL,
	patronymic TEXT,
	experience INT,
	birthday DATE NOT NULL,
	gender BOOL NOT NULL
);

CREATE TABLE sportsmen_coaches (
	sportsman_id UUID REFERENCES sportsmen(id) ON DELETE RESTRICT,
	coach_id UUID REFERENCES coaches(id) ON DELETE RESTRICT,
	CONSTRAINT "PK_sportsmen_coaches"
	PRIMARY KEY (sportsman_id, coach_id)
);

CREATE TABLE comp_accesses (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	sportsman_id UUID REFERENCES sportsmen(id) ON DELETE CASCADE,
	institution TEXT NOT NULL,
	validity DATE NOT NULL	
);

CREATE TABLE antidopings (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	sportsman_id UUID REFERENCES sportsmen(id) ON DELETE CASCADE,
	validity DATE NOT NULL
);

CREATE TABLE t_camps (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	city TEXT NOT NULL,
	address TEXT NOT NULL,
	beg_date DATE NOT NULL,
	end_date DATE NOT NULL CHECK (end_date >= beg_date)
);

CREATE TABLE competitions (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	city TEXT NOT NULL,
	address TEXT NOT NULL,
	beg_date DATE NOT NULL,
	end_date DATE NOT NULL CHECK (end_date >= beg_date),
	age TEXT NOT NULL,
	min_sports_category TEXT NOT NULL,
	antidoping BOOL NOT NULL
);

CREATE TABLE t_camp_applications (
	sportsman_id UUID REFERENCES sportsmen(id) ON DELETE RESTRICT,
	t_camp_id UUID REFERENCES t_camps(id) ON DELETE CASCADE,
	CONSTRAINT "PK_t_camp_applications"
	PRIMARY KEY (sportsman_id, t_camp_id)
);

CREATE TABLE comp_applications (
	sportsman_id UUID REFERENCES sportsmen(id) ON DELETE RESTRICT,
	competition_id UUID REFERENCES competitions(id) ON DELETE CASCADE,
	weight_category INT NOT NULL CHECK (weight_category > 0),
	start_snatch INT NOT NULL,
	start_clean_and_jerk INT NOT NULL,
	CONSTRAINT "PK_comp_applications"
	PRIMARY KEY (sportsman_id, competition_id)
);

CREATE TABLE results (
	sportsman_id UUID REFERENCES sportsmen(id) ON DELETE RESTRICT,
	competition_id UUID REFERENCES competitions(id) ON DELETE CASCADE,
	weight_category INT NOT NULL CHECK (weight_category > 0),
	snatch INT NOT NULL,
	clean_and_jerk INT NOT NULL,
	place INT NOT NULL,
	CONSTRAINT "PK_results"
	PRIMARY KEY (sportsman_id, competition_id)
);

CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	email TEXT NOT NULL CHECK (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
	password TEXT NOT NULL,
	role TEXT NOT NULL,
	role_id UUID
);

CREATE OR REPLACE FUNCTION get_upcoming_comps()
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


CREATE OR REPLACE FUNCTION get_upcoming_comps_sm(sm_id UUID) 
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


CREATE OR REPLACE FUNCTION get_upcoming_tcamps_sm(sm_id UUID) 
RETURNS TABLE (
    id UUID,
	city TEXT,
	address TEXT,
	beg_date DATE,
	end_date DATE
) AS $$
BEGIN
    RETURN QUERY
    SELECT tc.id, tc.city, tc.address, tc.beg_date, tc.end_date
    FROM t_camps tc
    JOIN t_camp_applications tca ON tc.id = tca.t_camp_id
    WHERE tca.sportsman_id = sm_id
    AND tc.beg_date > CURRENT_DATE;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION avg_exe_time(query_text TEXT, num_iterations INTEGER) RETURNS FLOAT AS $$
DECLARE
    total_time FLOAT := 0;
    i INTEGER := 0;
	beg FLOAT := 0;
	endd FLOAT := 0;
BEGIN
    WHILE i < num_iterations LOOP
		beg := EXTRACT(EPOCH FROM clock_timestamp());
        EXECUTE query_text;
		endd := EXTRACT(EPOCH FROM clock_timestamp());
        total_time := total_time + (endd - beg);
        i := i + 1;
    END LOOP;
    RETURN (total_time / num_iterations) * 1000;
END;
$$ LANGUAGE plpgsql;

/*INSERT INTO competitions (name, city, address, beg_date, end_date, age, min_sports_category, antidoping)
VALUES ('A', 'B', 'C', '2023-12-10', '2023-12-12', 'мужчины, женщины', 'I', true);*/

INSERT INTO sportsmen (id, surname, name, patronymic, birthday, moscow_team, sports_category, gender)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc1115', 'Иванов', 'Иван', 'Иванович', '2000-12-10', true, 'I', false);

/*INSERT INTO t_camps (id, city, address, beg_date, end_date)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc1110', 'И', 'И', '2024-12-10', '2024-12-12');*/

INSERT INTO users (email, password, role, role_id)
VALUES ('ivan@mail.ru', '123', 'sportsman', 'f8f26e6d-3e36-416a-9ff0-253876fc1115');

INSERT INTO competitions (id, name, city, address, beg_date, end_date, age, min_sports_category, antidoping)
VALUES ('a8f26e6d-3e36-416a-9ff0-253876fc1115', 'Кубок1', 'Москва', 'Улица 1', '2023-12-10', '2023-12-12', 'мужчины, женщины', 'I', false);

INSERT INTO competitions (id, name, city, address, beg_date, end_date, age, min_sports_category, antidoping)
VALUES ('b8f26e6d-3e36-416a-9ff0-253876fc1115', 'Кубок2', 'Москва', 'Улица 1', '2023-12-10', '2023-12-12', 'мужчины, женщины', 'I', false);

INSERT INTO t_camps (city, address, beg_date, end_date)
VALUES ('Москва', 'Улица 1', '2024-12-10', '2024-12-12');

INSERT INTO results (sportsman_id, competition_id, weight_category, snatch, clean_and_jerk, place)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc1115', 'b8f26e6d-3e36-416a-9ff0-253876fc1115', 81, 100, 120, 2);

INSERT INTO results (sportsman_id, competition_id, weight_category, snatch, clean_and_jerk, place)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc1115', 'a8f26e6d-3e36-416a-9ff0-253876fc1115', 87, 110, 150, 1);

INSERT INTO users (email, password, role)
VALUES ('sec@mail.ru', '123', 'secretary');

INSERT INTO users (email, password, role)
VALUES ('comp@mail.ru', '123', 'competition organizer');

INSERT INTO users (email, password, role)
VALUES ('tcamp@mail.ru', '123', 'training camp organizer');

INSERT INTO sportsmen (id, surname, name, patronymic, birthday, moscow_team, sports_category, gender)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc111c', 'Петрова', 'Полина', 'Петровна', '2000-09-10', true, 'КМС', true);

INSERT INTO coaches (surname, name, patronymic, birthday, experience, gender)
VALUES ('Семенов', 'Семен', 'Семенович', '1990-09-11', 10, false);

INSERT INTO coaches (id, surname, name, patronymic, birthday, experience, gender)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc1111', 'Сидоров', 'Алексей', 'Михайлович', '1980-09-11', 20, false);

INSERT INTO users (email, password, role, role_id)
VALUES ('alex@mail.ru', '123', 'coach', 'f8f26e6d-3e36-416a-9ff0-253876fc1111');

INSERT INTO sportsmen_coaches (sportsman_id, coach_id)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc1115', 'f8f26e6d-3e36-416a-9ff0-253876fc1111');

INSERT INTO sportsmen_coaches (sportsman_id, coach_id)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc111c', 'f8f26e6d-3e36-416a-9ff0-253876fc1111');

INSERT INTO comp_accesses (sportsman_id, institution, validity)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc111c', 'Диспансер 1', '2025-09-10');

INSERT INTO comp_accesses (sportsman_id, institution, validity)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc1115', 'Диспансер 2', '2025-09-10');

INSERT INTO antidopings (sportsman_id, validity)
VALUES ('f8f26e6d-3e36-416a-9ff0-253876fc111c', '2025-09-10');

COPY sportsmen
FROM '/app_mpa/sql/data/sportsmenData.csv' DELIMITER ',';

COPY comp_accesses
FROM '/app_mpa/sql/data/accessData.csv' DELIMITER ',';

COPY antidopings
FROM '/app_mpa/sql/data/antidopingData.csv' DELIMITER ',';

COPY coaches
FROM '/app_mpa/sql/data/coachesData.csv' DELIMITER ',';

COPY competitions
FROM '/app_mpa/sql/data/competitionsData.csv' DELIMITER ',';

COPY results
FROM '/app_mpa/sql/data/resultsData.csv' DELIMITER ',';

COPY sportsmen_coaches
FROM '/app_mpa/sql/data/sportsmenCoachesData.csv' DELIMITER ',';

COPY t_camps
FROM '/app_mpa/sql/data/t_campsData.csv' DELIMITER ',';