package domains

import (
	"time"

	"gorm.io/gorm"
)

type AppSystem struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"column:name;type:nvarchar(255);not null" json:"name"`
	Desc      string `gorm:"column:desc;type:nvarchar(max);not null" json:"description"` // Description in Thai
	Category  string `gorm:"column:category;type:varchar(50);not null" json:"category"`  // office, production, quality, support
	Href      string `gorm:"column:href;type:nvarchar(max);not null" json:"href"`
	Icon      string `gorm:"column:icon;type:nvarchar(max)" json:"icon"`             // React icon name or base64 encoded icon
	ParentID  *int   `gorm:"column:parent_id;index;default:null" json:"parent_id"`   // null = root menu item
	SortOrder int    `gorm:"column:sort_order;default:0;not null" json:"sort_order"` // ลำดับแสดงผล (0=ก่อนสุด)

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

type AppSystemMenu struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Category    string          `json:"category,omitempty"`
	Href        string          `json:"href,omitempty"`
	SortOrder   int             `json:"sort_order"`
	Children    []AppSystemMenu `json:"children,omitempty"`
}

// BuildAppSystemTree แปลง flat list เป็น hierarchical tree structure
func BuildAppSystemTree(systems []AppSystem) []AppSystemMenu {
	// จัดกลุ่ม children ตาม parent_id
	childrenMap := make(map[int][]AppSystem) // map of parentID -> children
	rootItems := []AppSystem{}               // items with no parent

	for _, sys := range systems {
		if sys.ParentID == nil {
			rootItems = append(rootItems, sys)
		} else {
			childrenMap[*sys.ParentID] = append(childrenMap[*sys.ParentID], sys)
		}
	}

	// Sort children ตาม SortOrder
	sortByOrder := func(items []AppSystem) {
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
	var buildTree func(parentID int) []AppSystemMenu
	buildTree = func(parentID int) []AppSystemMenu {
		var menus []AppSystemMenu
		children := childrenMap[parentID]
		sortByOrder(children)

		for _, sys := range children {
			menu := AppSystemMenu{
				ID:          sys.ID,
				Name:        sys.Name,
				Description: sys.Desc,
				Category:    sys.Category,
				Href:        sys.Href,
				SortOrder:   sys.SortOrder,
				Children:    buildTree(sys.ID), // Recursive
			}
			menus = append(menus, menu)
		}

		return menus
	}

	// Build tree for root items
	var result []AppSystemMenu
	for _, root := range rootItems {
		menu := AppSystemMenu{
			ID:          root.ID,
			Name:        root.Name,
			Description: root.Desc,
			Category:    root.Category,
			Href:        root.Href,
			SortOrder:   root.SortOrder,
			Children:    buildTree(root.ID),
		}
		result = append(result, menu)
	}

	return result
}

// TableName specifies the table name for the AppSystem model
func (AppSystem) TableName() string {
	return "ps_app_systems"
}
