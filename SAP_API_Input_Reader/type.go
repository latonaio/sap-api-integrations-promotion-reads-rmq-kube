package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	PromotionCollection struct {
		ObjectID                           string `json:"ObjectID"`
		ETag                               string `json:"ETag"`
		AccountType                        string `json:"AccountType"`
		AccountTypeText                    string `json:"AccountTypeText"`
		Currency                           string `json:"Currency"`
		CurrencyText                       string `json:"CurrencyText"`
		ID                                 string `json:"ID"`
		Name                               string `json:"Name"`
		Objective                          string `json:"Objective"`
		ObjectiveText                      string `json:"ObjectiveText"`
		Priority                           string `json:"Priority"`
		PriorityText                       string `json:"PriorityText"`
		ProductPlanningBasisCode           string `json:"ProductPlanningBasisCode"`
		ProductPlanningBasisCodeText       string `json:"ProductPlanningBasisCodeText"`
		PromotionType                      string `json:"PromotionType"`
		PromotionTypeText                  string `json:"PromotionTypeText"`
		ConsistencyStatusCode              string `json:"ConsistencyStatusCode"`
		ConsistencyStatusCodeText          string `json:"ConsistencyStatusCodeText"`
		ExternalStatusCode                 string `json:"ExternalStatusCode"`
		ExternalStatusCodeText             string `json:"ExternalStatusCodeText"`
		LifeCycleStatusCode                string `json:"LifeCycleStatusCode"`
		LifeCycleStatusCodeText            string `json:"LifeCycleStatusCodeText"`
		CreationDate                       string `json:"CreationDate"`
		LastChangeDate                     string `json:"LastChangeDate"`
		Tactic                             string `json:"Tactic"`
		TacticText                         string `json:"TacticText"`
		PlannigAccountName                 string `json:"PlannigAccountName"`
		PlanningAccountID                  string `json:"PlanningAccountID"`
		PlanPeriodStartDate                string `json:"PlanPeriodStartDate"`
		PlanPeriodEndDate                  string `json:"PlanPeriodEndDate"`
		BuyingPeriodStartDate              string `json:"BuyingPeriodStartDate"`
		BuyingPeriodEndDate                string `json:"BuyingPeriodEndDate"`
		EmployeeResponsible                string `json:"EmployeeResponsible"`
		EmployeeResposibleID               string `json:"EmployeeResposibleID"`
		SalesUnitOrganisationCentreID      string `json:"SalesUnitOrganisationCentreID"`
		SalesUnitOrganisationCentreName    string `json:"SalesUnitOrganisationCentreName"`
		SalesOrganisationID                string `json:"SalesOrganisationID"`
		SalesOrganisationName              string `json:"SalesOrganisationName"`
		DistributionChannelCode            string `json:"DistributionChannelCode"`
		DistributionChannelCodeText        string `json:"DistributionChannelCodeText"`
		DivisionCode                       string `json:"DivisionCode"`
		DivisionCodeText                   string `json:"DivisionCodeText"`
		SalesTerritoryID                   string `json:"SalesTerritoryID"`
		SalesTerritoryName                 string `json:"SalesTerritoryName"`
		ActualPeriodStartDate              string `json:"ActualPeriodStartDate"`
		ActualPeriodEndDate                string `json:"ActualPeriodEndDate"`
		Campaign                           string `json:"Campaign"`
		CampaignDescription                string `json:"CampaignDescription"`
		TargetGroupID                      string `json:"TargetGroupID"`
		TargetGroupDescription             string `json:"TargetGroupDescription"`
		ExternalID                         string `json:"ExternalID"`
		InformationLifeCycleStatusCode     string `json:"InformationLifeCycleStatusCode"`
		InformationLifeCycleStatusCodeText string `json:"InformationLifeCycleStatusCodeText"`
		Cancel                             bool   `json:"Cancel"`
		RevokeCancellation                 bool   `json:"RevokeCancellation"`
		EntityLastChangedOn                string `json:"EntityLastChangedOn"`
		PromotionParty                     struct {
			ObjectID       string `json:"ObjectID"`
			ParentObjectID string `json:"ParentObjectID"`
			ETag           string `json:"ETag"`
			ID             string `json:"ID"`
			RoleCode       string `json:"RoleCode"`
			RoleCodeText   string `json:"RoleCodeText"`
			Name           string `json:"Name"`
			MainIndicator  bool   `json:"MainIndicator"`
		} `json:"PromotionParty"`
	} `json:"PromotionCollection"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	PromotionCode string   `json:"promotion_code"`
	Deleted       bool     `json:"deleted"`
}