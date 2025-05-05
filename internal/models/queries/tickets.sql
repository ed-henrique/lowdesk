-- name: GetAllTickets :many
SELECT * FROM TICKETS;

-- name: InsertTicket :one
INSERT INTO TICKETS
    (title, content)
VALUES
    (@title, @content)
RETURNING ID;
