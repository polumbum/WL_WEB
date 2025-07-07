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



SELECT * FROM get_upcoming_comps();
SELECT * FROM get_upcoming_tcamps();
SELECT * FROM get_upcoming_comps_sm('f8f26e6d-3e36-416a-9ff0-253876fc1115');
SELECT * FROM get_upcoming_comps_sm('f8f26e6d-3e36-416a-9ff0-253876fc1115');