package account

import "time"

type EstimationHeader struct {
	ID                int       `json:"id,omitempty" db:"ID"`
	CustomerID        int       `json:"customerid,omitempty" db:"CustomerID"`
	EffectiveDate     time.Time `json:"effectiveDate,omitempty" db:"EffectiveDate"`
	Type              string    `json:"type,omitempty" db:"Type"`
	Size              string    `json:"size,omitempty" db:"Size"`
	Disabled          int       `json:"disabled,omitempty" db:"Disabled"`
	Remarks           string    `json:"remarks,omitempty" db:"Remarks"`
	CreatedBy         string    `json:"createdby,omitempty" db:"CreatedBy"`
	CreateDate        time.Time `json:"createdate,omitempty" db:"CreateDate"`
	LastModBy         string    `json:"lastmodby,omitempty" db:"LastModBy"`
	LastMod           time.Time `json:"lastmod,omitempty" db:"LastMod"`
	Purpose           string    `json:"purpose,omitempty" db:"Purpose"`
	EstimationCounter int       `json:"estimationcounter,omitempty" db:"EstimationCounter"`
	RevIncTax         float64   `json:"revInctax,omitempty" db:"RevIncTax"`
	Approved          int       `json:"approved,omitempty" db:"Approved"`
	ApprovedDateTime  time.Time `json:"approveddatetime,omitempty" db:"ApprovedDateTime"`
	ApprovedBy        string    `json:"approvedby,omitempty" db:"ApprovedBy"`
	EstimationNo      string    `json:"estimationno,omitempty" db:"EstimationNo"`
}

type EstimationDetails struct {
	EstimationID       int     `json:"estimationid,omitempty" db:"EstimationID"`
	EstimationDetailID int     `json:"estimationDetailID,omitempty" db:"EstimationDetailID"`
	IsExpense          int     `json:"isexpense,omitempty" db:"IsExpense"`
	ItemID             int     `json:"itemid,omitempty" db:"ItemID"`
	UOM                int     `json:"uom,omitempty" db:"UOM"`
	Amount             float64 `json:"amount,omitempty" db:"Amount"`
	Remarks            string  `json:"remarks,omitempty" db:"Remarks"`
	PerDoc             int     `json:"perdoc,omitempty" db:"PerDoc"`
	Jalur              string  `json:"jalur,omitempty" db:"Jalur"`
	AreaID             int     `json:"areaid,omitempty" db:"AreaID"`
	PortID             int     `json:"portid,omitempty" db:"PortID"`
	CSize              string  `json:"csize,omitempty" db:"CSize"`
	DestID             int     `json:"destid,omitempty" db:"DestID"`
	ByQty              float64 `json:"byqty,omitempty" db:"ByQty"`
	SupplierID         int     `json:"supplierid,omitempty" db:"SupplierID"`
	CurrID             int     `json:"currid,omitempty" db:"CurrID"`
	ShippingLineID     int     `json:"shippinglineid,omitempty" db:"ShippingLineID"`
	JenisKarantina     string  `json:"jeniskarantina,omitempty" db:"JenisKarantina"`
	SandaranID         int     `json:"sandaranid,omitempty" db:"SandaranID"`
}

type Estimation struct {
	EstHeader  EstimationHeader    `json:"estHeader,omitempty"`
	EstDetails []EstimationDetails `json:"estDetail,omitempty"`
}

var (
	Tbl_estimation = map[string]string{
		"header": "est_estimation",
		"detail": "est_estimationdetails",
	}

	EstColumn = map[string]string{
		"header": "CustomerID, EffectiveDate, Type, Size, Disabled, Remarks, CreatedBy, CreateDate, LastModBy, LastMod, Purpose, EstimationCounter, RevIncTax, Approved, ApprovedBy, ApprovedDateTime, EstimationNo",
		"detail": "EstimationID, EstimationDetailID, IsExpense, ItemID, UOM, Amount, Remarks, PerDoc, Jalur, AreaID, PortID, CSize, DestID, ByQty, SupplierID, CurrID, ShippingLineID, JenisKarantina, SandaranID",
	}

	EstValue = map[string]string{
		"header": ":CustomerID, :EffectiveDate, :Type, :Size, :Disabled, :Remarks, :CreatedBy, :CreateDate, :LastModBy, :LastMod, :Purpose, :EstimationCounter, :RevIncTax, :Approved, :ApprovedBy, :ApprovedDateTime, :EstimationNo",
		"detail": ":EstimationID, :EstimationDetailID, :IsExpense, :ItemID, :UOM, :Amount, :Remarks, :PerDoc, :Jalur, :AreaID, :PortID, :CSize, :DestID, :ByQty, :SupplierID, :CurrID, :ShippingLineID, :JenisKarantina, :SandaranID",
	}

	EstColVal = map[string]string{
		"header": "CustomerID =:CustomerID, EffectiveDate =:EffectiveDate, Type = :Type, Size = :Size, Disabled = :Disabled, Remarks = :Remarks, CreatedBy = :CreatedBy, CreateDate = :CreateDate, LastModBy = :LastModBy, LastMod = :LastMod, Purpose = :Purpose, EstimationCounter = :EstimationCounter, RevIncTax = :RevIncTax, Approved = :Approved, ApprovedBy = :ApprovedBy, EstimationNo = :EstimationNo",
		"detail": "IsExpense = :IsExpense, EstimationDetailID := EstimationDetailID, ItemID = :ItemID, UOM = :UOM, Amount = :Amount, Remarks = :Remarks, PerDoc = :PerDoc, Jalur = :Jalur, AreaID = :AreaID, PortID = :PortID, CSize = :CSize, DestID = :DestID, ByQty = :ByQty, SupplierID = :SupplierID, CurrID = :CurrID, ShippingLineID = :ShippingLineID, JenisKarantina = :JenisKarantina, SandaranID = :SandaranID",
	}

	EstList []EstimationHeader
)

type Services interface {
	FindByID(id string) *Estimation
	List() *Estimation
	Update(id string, mdl Estimation) error
	Delete(id string) error
	Create(data Estimation) error
}
