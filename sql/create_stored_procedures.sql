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
    VALUES (in_user_id, p_movie_id, in_rating, in_review_text);
END;
$$;

-- Stored procedure to update a user review
CREATE OR REPLACE PROCEDURE update_user_review(
    IN in_user_id BIGINT,
    IN in_movie_id BIGINT,
    IN in_rating INT,
    IN in_review_text TEXT
)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE user_reviews
    SET rating = in_rating,
        review_text = in_review_text,
        updated_at = CURRENT_TIMESTAMP
    WHERE user_id = in_user_id AND movie_id = in_movie_id;
END;
$$;

-- Stored procedure to delete a user review
CREATE OR REPLACE PROCEDURE delete_user_review(
    IN in_user_id BIGINT,
    IN in_movie_id BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    DELETE FROM user_reviews
    WHERE user_id = in_user_id AND movie_id = in_movie_id;
END;
$$;

