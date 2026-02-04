package domains

import (
	"time"

	"gorm.io/gorm"
)

type CustomerManual struct {
	CustomerManualID   int            `gorm:"primaryKey;autoIncrement"`
	CustomerManualName string         `gorm:"type:nvarchar(255);not null"`
	Desc               string         `gorm:"type:nvarchar(max);not null"`
	Category           string         `gorm:"type:varchar(50);not null"`
	FileName           string         `gorm:"type:nvarchar(max);not null"`
	ParentID           *int           `gorm:"column:parent_id;index;default:null"`
	SortOrder          int            `gorm:"column:sort_order;default:0;not null"`
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

type CustomerManualMenu struct {
	ID        int                  `json:"id"`
	Name      string               `json:"name"`
	Desc      string               `json:"desc,omitempty"`
	Category  string               `json:"category,omitempty"`
	FileName  string               `json:"file_name,omitempty"`
	SortOrder int                  `json:"sort_order"`
	Children  []CustomerManualMenu `json:"children,omitempty"`
}

// BuildCustomerManualTree แปลง flat list เป็น hierarchical tree structure
func BuildCustomerManualTree(manuals []CustomerManual) []CustomerManualMenu {
	// จัดกลุ่ม children ตาม parent_id
	childrenMap := make(map[int][]CustomerManual) // map of parentID -> children
	rootItems := []CustomerManual{}               // items with no parent

	for _, manual := range manuals {
		if manual.ParentID == nil {
			rootItems = append(rootItems, manual)
		} else {
			childrenMap[*manual.ParentID] = append(childrenMap[*manual.ParentID], manual)
		}
	}

	// Sort children ตาม SortOrder
	sortByOrder := func(items []CustomerManual) {
		for i := 0; i < len(items); i++ {
			for j := i + 1; j < len(items); j++ {
				if items[i].SortOrder > items[j].SortOrder {
					items[i], items[j] = items[j], items[i]
				}
			}
		}
	}

	// Sort root items
	sortByOrder(rootItems)

	// Recursive function สร้าง menu tree
	var buildTree func(parentID int) []CustomerManualMenu
	buildTree = func(parentID int) []CustomerManualMenu {
		var menus []CustomerManualMenu
		children := childrenMap[parentID]
		sortByOrder(children)

		for _, manual := range children {
			menu := CustomerManualMenu{
				ID:        manual.CustomerManualID,
				Name:      manual.CustomerManualName,
				Desc:      manual.Desc,
				Category:  manual.Category,
				FileName:  manual.FileName,
				SortOrder: manual.SortOrder,
				Children:  buildTree(manual.CustomerManualID), // Recursive
			}
			menus = append(menus, menu)
		}

		return menus
	}

	// Build tree for root items
	var result []CustomerManualMenu
	for _, root := range rootItems {
		menu := CustomerManualMenu{
			ID:        root.CustomerManualID,
			Name:      root.CustomerManualName,
			Desc:      root.Desc,
			Category:  root.Category,
			FileName:  root.FileName,
			SortOrder: root.SortOrder,
			Children:  buildTree(root.CustomerManualID),
		}
		result = append(result, menu)
	}

	return result
}

func (CustomerManual) TableName() string {
	return "customer_manuals"
}
