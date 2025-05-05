package services

import (
	"cafeweb-backend/dto"
	"cafeweb-backend/models"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type ProductService struct {
	db *gorm.DB
}

// NewProductService - ฟังก์ชันสำหรับสร้าง ProductService
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		db: db,
	}
}

func (s *ProductService) CreateProduct(req dto.ProductRequest) (out dto.ProductResponse, err error) {
	tx := s.db.Begin()
	fmt.Println("req", req)

	product := models.Products{
		ProductName: req.ProductName,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
		CategoryID:  req.CategoryID,
		CafeID:      req.CafeID,
		Status:      models.StatusType(req.Status),
		Description: req.Description,
	}

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return out, err
	}

	if err := tx.Preload("Category").Where("id = ?", product.ID).First(&product).Error; err != nil {
		tx.Rollback()
		return out, err
	}

	copier.Copy(&out, &product)
	if product.Category.CategoryName != "" {
		out.CategoryName = product.Category.CategoryName
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return out, err
	}
	return out, nil
}

func (s *ProductService) GetProduct(cafeID string) (out []dto.ProductResponse, err error) {
	tx := s.db.Begin()
	var products []models.Products

	fmt.Println("cafeID :: ", cafeID)

	if err := tx.Preload("Category").Where("cafe_id = ?", cafeID).
		Find(&products).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	fmt.Println("cafeID products:: ", products)

	if len(products) == 0 {
		return nil, errors.New("no products found for this cafe")
	} else {
		fmt.Println("Found products:", products)
	}

	for _, p := range products {
		var pr dto.ProductResponse
		copier.Copy(&pr, &p)
		pr.CategoryName = p.Category.CategoryName
		out = append(out, pr)
	}

	tx.Commit()
	return out, nil
}

func (s *ProductService) GetProductByID(cafeID string, ID string) (out dto.ProductResponse, err error) {
	tx := s.db.Begin()
	var product models.Products

	fmt.Println("cafeID :: ", cafeID)

	if err := tx.Preload("Category").Where("cafe_id = ? and id = ?", cafeID, ID).
		Find(&product).Error; err != nil {
		tx.Rollback()
		return out, err
	}

	copier.Copy(&out, &product)
	if product.Category.CategoryName != "" {
		out.CategoryName = product.Category.CategoryName
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return out, err
	}
	return out, nil
}

func (s *ProductService) UpdateProduct(cafeID string, id string, req dto.ProductRequest) (out dto.ProductResponse, err error) {
	tx := s.db.Begin()

	var product models.Products
	if err := tx.First(&product, "cafe_id = ? and id = ?", cafeID, id).Error; err != nil {
		tx.Rollback()
		return out, fmt.Errorf("product not found: %w", err)
	}

	product.ProductName = req.ProductName
	product.Price = req.Price
	product.ImageURL = req.ImageURL
	product.CategoryID = req.CategoryID
	product.CafeID = req.CafeID
	product.Status = models.StatusType(req.Status)
	product.Description = req.Description

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return out, fmt.Errorf("failed to update product: %w", err)
	}

	if err := tx.Preload("Category").Where("id = ?", product.ID).First(&product).Error; err != nil {
		tx.Rollback()
		return out, fmt.Errorf("failed to fetch updated product: %w", err)
	}

	copier.Copy(&out, &product)
	if product.Category.CategoryName != "" {
		out.CategoryName = product.Category.CategoryName
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return out, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return out, nil
}

func (s *ProductService) DeleteMultipleProducts(ids []string) error {
	if len(ids) == 0 {
		return errors.New("no product IDs provided")
	}

	tx := s.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Where("id IN (?)", ids).Delete(&models.Products{})

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("no products were deleted")
	}

	return tx.Commit().Error
}

func (s *ProductService) CreateCategories(req dto.CategoryRequest) (err error) {
	tx := s.db.Begin()

	var existingCategory models.Categories
	if err := tx.Where("category_name = ? AND cafe_id = ?", req.CategoryName, req.CafeID).
		First(&existingCategory).Error; err == nil {

		tx.Rollback()
		return errors.New("category already exists for this cafe")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {

		tx.Rollback()
		return err
	}

	category := models.Categories{
		CategoryName: req.CategoryName,
		CafeID:       req.CafeID,
	}
	if err := tx.Create(&category).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *ProductService) GetCategories(cafeID string) (out []dto.CategoryResponse, err error) {
	tx := s.db.Begin()

	var category []models.Categories
	if err := tx.Where("cafe_id =?", cafeID).Find(&category).Error; err != nil {
		tx.Rollback()
		return out, nil
	}
	if len(category) == 0 {
		return nil, errors.New("no category found for this cafe")
	} else {
		fmt.Println("Found category:", category)
	}

	for _, p := range category {
		var pr dto.CategoryResponse
		copier.Copy(&pr, &p)
		out = append(out, pr)
	}
	tx.Commit()
	return out, nil
}
