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

	// ใช้ tx เพื่อทำการสร้าง product ภายใน transaction
	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback() // ย้อนกลับหากมีข้อผิดพลาด
		return out, err
	}

	// ถ้าไม่มีข้อผิดพลาด ให้ยืนยันการทำธุรกรรม
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // ย้อนกลับหาก Commit ล้มเหลว
		return out, err
	}

	out.CafeID = product.CafeID
	out.CategoryID = product.CategoryID
	out.ImageURL = product.ImageURL
	out.Price = product.Price
	out.ProductName = product.ProductName
	out.Status = product.Status
	out.Description = product.Description

	// คืนค่าข้อมูล product ที่ถูกสร้างแล้ว
	return out, nil
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
