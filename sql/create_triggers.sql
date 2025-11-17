-- This SQL script creates triggers to automatically update the 'updated_at' timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger for user_reviews table
CREATE TRIGGER set_user_reviews_updated_at
BEFORE UPDATE ON user_reviews
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();