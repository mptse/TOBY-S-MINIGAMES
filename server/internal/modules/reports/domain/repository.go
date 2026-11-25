package domain

import "context"

// ReportRepository define las operaciones de persistencia para los reportes de incidencias.
type ReportRepository interface {
	Create(ctx context.Context, report *Report) error
	ListByStatus(ctx context.Context, status ReportStatus, limit, offset int) ([]*Report, error)
}
