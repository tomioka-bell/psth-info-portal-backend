package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"backend/internal/core/domains"
	ports "backend/internal/core/ports/repositories"
)

type ProductRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ports.ProductRepository {
	if err := db.AutoMigrate(&domains.Product{}); err != nil {
		fmt.Printf("failed to auto migrate: %v", err)
	}
	return &ProductRepositoryDB{db: db}
}

func (r *ProductRepositoryDB) CreateProduct(n *domains.Product) error {
	if err := r.db.Create(n).Error; err != nil {
		fmt.Printf("CreateUserRepository error: %v\n", err)
		return err
	}
	return nil
}

func (r *ProductRepositoryDB) GetAllProducts(limit int) ([]domains.Product, error) {
	var products []domains.Product
	if err := r.db.Limit(limit).Order("created_at DESC").Find(&products).Error; err != nil {
		fmt.Printf("GetAllProducts error: %v\n", err)
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryDB) GetProductByID(productID string) (*domains.Product, error) {
	var product domains.Product
	if err := r.db.Where("product_id = ?", productID).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("product not found")
		}
		fmt.Printf("GetProductByID error: %v\n", err)
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryDB) GetRecommendedProducts(limit int) ([]domains.Product, error) {
	var products []domains.Product

	if err := r.db.
		Where("recommend = ?", true).
		Order("created_at DESC").
		Limit(limit).
		Find(&products).Error; err != nil {

		fmt.Printf("GetRecommendedProducts error: %v\n", err)
		return nil, err
	}

	return products, nil
}

func (r *ProductRepositoryDB) UpdateProduct(productID string, product *domains.Product) error {
	if err := r.db.Model(&domains.Product{}).Where("product_id = ?", productID).Select("*").Updates(product).Error; err != nil {
		fmt.Printf("UpdateProduct error: %v\n", err)
		return err
	}
	return nil
}

func (r *ProductRepositoryDB) DeleteProduct(productID string) error {
	if err := r.db.Where("product_id = ?", productID).Delete(&domains.Product{}).Error; err != nil {
		fmt.Printf("DeleteProduct error: %v\n", err)
		return err
	}
	return nil
}

func (r *ProductRepositoryDB) SearchProductsByName(keyword string, limit int) ([]domains.Product, error) {

	var products []domains.Product

	if err := r.db.Debug().
		Where("UPPER(product_name) LIKE UPPER(?)", "%"+keyword+"%").
		Order("created_at DESC").
		Limit(limit).
		Find(&products).Error; err != nil {

		fmt.Printf("SearchProductsByName error: %v\n", err)
		return nil, err
	}

	return products, nil
}
