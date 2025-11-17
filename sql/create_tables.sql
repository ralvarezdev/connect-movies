-- user_reviews table
CREATE TABLE IF NOT EXISTS user_reviews (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    movie_id BIGINT NOT NULL,
    rating INT CHECK (rating >= 1 AND rating <= 10),
    review_text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
);
CREATE UNIQUE INDEX IF NOT EXISTS user_reviews_unique_user_movie_review
ON user_reviews (user_id, movie_id);
