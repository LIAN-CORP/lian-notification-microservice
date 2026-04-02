-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE notification_events(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_type VARCHAR(255) NOT NULL,
    debt_id UUID NOT NULL,
    client_id UUID NOT NULL,
    payload JSONB NOT NULL,
    received_at TIMESTAMP NOT NULL DEFAULT NOW(),
    processed BOOLEAN DEFAULT FALSE
);

CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id UUID NOT NULL,
    client_id UUID NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    message TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    provider VARCHAR(50) DEFAULT 'AWS_SNS',
    provider_response TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    snet_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_notification_event FOREIGN KEY (event_id) REFERENCES notification_events(id) ON DELETE CASCADE
);

CREATE TABLE notification_attemps (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    notification_id UUID NOT NULL,
    attempt_number INT NOT NULL,
    status VARCHAR(20) NOT NULL,
    error_message TEXT,
    attempt_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_notification_attempt FOREIGN KEY (notification_id) REFERENCES notifications(id) ON DELETE CASCADE
);

CREATE TABLE debt_snapshot (
    debt_id UUID PRIMARY KEY,
    client_id UUID NOT NULL,
    total_amount NUMERIC(15,2),
    remaining_amount NUMERIC(15,2),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE notification_templates(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_type VARCHAR(255) NOT NULL,
    template TEXT NOT NULL
);

-- ============== INDICES ==============
CREATE INDEX idx_events_debt_id ON notification_events(debt_id);
CREATE INDEX idx_events_client_id ON notification_events(client_id);
CREATE INDEX idx_events_processed ON notification_events(processed);
CREATE INDEX idx_notifications_status ON notifications(status);
CREATE INDEX idx_notifications_client_id ON notifications(client_id);
CREATE INDEX idx_notifications_event_id ON notifications(event_id);
CREATE INDEX idx_attempts_notification_id ON notification_attemps (notification_id);
CREATE INDEX idx_debt_snapshot_customer ON debt_snapshot(client_id);

-- +goose Down
DROP TABLE IF EXISTS notification_templates;
DROP TABLE IF EXISTS debt_snapshot;
DROP TABLE IF EXISTS notification_attemps;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS notification_events;