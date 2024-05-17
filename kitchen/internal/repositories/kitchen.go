package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"kitchenService/internal/config"
	"kitchenService/internal/models"
	"kitchenService/internal/ports"
)

type kitchenRepository struct {
	conn *pgx.Conn
}

func NewKitchenRepository(cfg *config.Config) (ports.KitchenRepository, error) {
	connectionString := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	fmt.Println(connectionString)
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	return &kitchenRepository{conn: conn}, nil
}

func (k *kitchenRepository) ProcessOrder(ctx context.Context, order *models.Order) (string, *models.OrderStatus, error) {
	tx, err := k.conn.Begin(ctx)
	if err != nil {
		return "", nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`INSERT INTO orders(id, status) VALUES ($1, $2)`,
		order.ID, order.Status)
	if err != nil {
		return "", nil, fmt.Errorf("failed to insert order: %v", err)
	}

	for _, item := range order.Items {
		_, err := tx.Exec(ctx,
			`INSERT INTO items(order_id, name, price, weight, created_at, updated_at, deleted_at, is_active)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			order.ID, item.Name, item.Price, item.Weight, item.CreatedAt, item.UpdatedAt, item.DeletedAt, item.IsActive)
		if err != nil {
			return "", nil, fmt.Errorf("failed to insert item: %v", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	orderStatus := models.Order{Status: models.OrderStatusPending}.Status

	return order.ID, &orderStatus, nil
}

func (k *kitchenRepository) ChangeOrderStatus(ctx context.Context, orderId string, status *models.OrderStatus) error {
	_, err := k.conn.Exec(ctx,
		`UPDATE orders SET status=$1 WHERE id=$2`,
		status, orderId)
	if err != nil {
		return fmt.Errorf("failed to change order status: %v", err)
	}

	return nil
}
