-- Обновите up миграцию
CREATE TABLE IF NOT EXISTS chat_rooms (
    id VARCHAR(255) PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    created_by VARCHAR(255),  -- Убрали NOT NULL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chat_message (
    id VARCHAR(255) PRIMARY KEY,
    type INTEGER NOT NULL,
    room_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    username TEXT NOT NULL,
    content TEXT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_chat_message_room FOREIGN KEY (room_id) REFERENCES chat_rooms(id)
);
CREATE INDEX IF NOT EXISTS idx_chat_message_room_id ON chat_message(room_id);
CREATE INDEX IF NOT EXISTS idx_chat_message_timestamp ON chat_message(timestamp);
CREATE INDEX IF NOT EXISTS idx_chat_message_user_id ON chat_message(user_id);