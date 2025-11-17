-- Stored procedure to create a user review
CREATE OR REPLACE PROCEDURE create_user_review(
    IN in_user_id BIGINT,
    IN in_movie_id BIGINT,
    IN in_rating INT,
    IN in_review_text TEXT
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO user_reviews (user_id, movie_id, rating, review_text)
    VALUES (in_user_id, in_movie_id, in_rating, in_review_text);
END;
$$;

-- Stored procedure to check if the user review exists
CREATE OR REPLACE PROCEDURE user_review_exists(
    IN in_user_id BIGINT,
    IN in_movie_id BIGINT,
    OUT out_user_review_found
)
LANGUAGE plpgsql
AS $$
BEGIN
    -- Check if the review exists and is not deleted
    SELECT EXISTS (
        SELECT 1 FROM user_reviews
        WHERE user_id = in_user_id AND movie_id = in_movie_id
        AND deleted_at IS NULL
    ) INTO out_user_review_found;
END;
$$;

-- Stored procedure to update a user review
CREATE OR REPLACE PROCEDURE update_user_review(
    IN in_user_id BIGINT,
    IN in_movie_id BIGINT,
    IN in_rating INT,
    IN in_review_text TEXT,
    OUT out_user_review_found BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
    -- Check if the user review exists
    CALL user_review_exists(in_user_id, in_movie_id, out_user_review_found);

    -- If the review exists, update it
    IF NOT out_user_review_found THEN
        RETURN;
    END IF;
    
    UPDATE user_reviews
    SET rating = in_rating,
        review_text = in_review_text
    WHERE user_id = in_user_id AND movie_id = in_movie_id
    AND deleted_at IS NULL;
END;
$$;

-- Stored procedure to delete a user review
CREATE OR REPLACE PROCEDURE delete_user_review(
    IN in_user_id BIGINT,
    IN in_movie_id BIGINT,
    OUT out_user_review_found
)
LANGUAGE plpgsql
AS $$
BEGIN
    -- Check if the user review exists
    CALL user_review_exists(in_user_id, in_movie_id, out_user_review_found);
    
    -- If the review exists, update it
    IF NOT out_user_review_found THEN
        RETURN;
    END IF;
    
    -- Update the deleted_at timestamp to mark as deleted
    UPDATE user_reviews
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE user_id = in_user_id AND movie_id = in_movie_id
    AND deleted_at IS NULL;
END;
$$;

-- Stored procedure to get a user review
CREATE OR REPLACE PROCEDURE get_user_review(
    IN in_user_id BIGINT,
    IN in_movie_id BIGINT,
    OUT out_rating INT,
    OUT out_review_text TEXT,
    OUT out_created_at TIMESTAMP,
    OUT out_updated_at TIMESTAMP,
    OUT out_user_review_found BOOLEAN
)
LANGUAGE plpgsql
AS $$
BEGIN
    -- Check if the user review exists
    CALL user_review_exists(in_user_id, in_movie_id, out_user_review_found);
    
    -- If the review does not exist, return
    IF NOT out_user_review_found THEN
        RETURN;
    END IF;
    
    -- Retrieve the review details
    SELECT rating, review_text, created_at, updated_at
    INTO out_rating, out_review_text, out_created_at, out_updated_at
    FROM user_reviews
    WHERE user_id = in_user_id AND movie_id = in_movie_id
    AND deleted_at IS NULL;
END;
$$;