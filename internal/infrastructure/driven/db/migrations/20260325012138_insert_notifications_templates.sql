-- +goose Up
INSERT INTO notification_templates (event_type, template) VALUES
('DEBT_CREATED', 'Hola {{clientName}}, adquiriste una deuda por ${{totalAmount}}. ¡Gracias por tu compra!'),
('DEBT_REMINDER', 'Hola {{clientName}}, recuerda que tienes una deuda activa por ${{totalAmount}}. Pagala lo más pronto posible para no dañar la reputación con el vendedor!!'),
('DEBT_INCREASE', 'Hola {{clientName}}, agregaste {{product.name}} por {{product.price}}. Tu deuda ahora es de ${{newTotalDebt}}'),
('PAYMENT_MADE', 'Hola {{clientName}}, recibimos tu pago de ${{paymentAmount}}. Tu saldo pendiente es de ${{remainingDebt}}. ¡Gracias!'),
('DEBT_FULL_PAID', 'Hola {{clientName}}, has pagado tu deuda en su totalidad. !Gracias por cumplir con tus pagos!');

-- +goose Down
DELETE FROM notification_templates;
