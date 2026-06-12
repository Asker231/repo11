package dictionary

type Config struct {
    ProductNumber           map[string]map[string]string `json:"productNumber"`
    Currency                map[string]string            `json:"currency"`
    CalculationMethod       map[string]string            `json:"calculationMethod"`
    CapitalizationType      map[string]string            `json:"capitalizationType"`
    FilialCode              map[string]string            `json:"filialCode"`
    TaskStatus              map[string]string            `json:"taskStatus"`
    Operation               Operation                     `json:"operation"`
    AccountCodes            AccountCodes                  `json:"accountCodes"`
    TaskAction              map[string]string            `json:"taskAction"`
    OperationAvailable      map[string]string            `json:"operationAvailable"`
    EventOperationType      EventOperationType            `json:"eventOperationType"`
    DealStatus              DealStatus                    `json:"dealStatus"`
    MidasType               []MidasType                   `json:"midasType"`
    CoverageTransitAccount  CoverageTransitAccount        `json:"coverageTransitAccount"`
    ProcessName             map[string]string            `json:"processName"`
    MovementType            map[string]string            `json:"movementType"`
    SystemName              map[string]string            `json:"systemName"`
    BaseRateList            BaseRateList                  `json:"base_rate_list"`
    DepositReleaseVersion   string                        `json:"depositReleaseVersion"`
}

type Operation struct {
    Create       string `json:"create"`
    Withdrawal   string `json:"withdrawal"`
    Replenishment string `json:"replenishment"`
    Delete       string `json:"delete"`
    TermUpdate   string `json:"term_update"`
    RateUpdate   string `json:"rate_update"`
    Terminate    Terminate `json:"terminate"`
}

type Terminate struct {
    Coverage string `json:"coverage"`
    Deposit  string `json:"deposit"`
}

type AccountCodes struct {
    CorpAccountCodes   string `json:"corpAccountCodes"`
    RetailAccountCodes string `json:"retailAccountCodes"`
}

type EventOperationType struct {
    AmountChange string   `json:"amountChange"`
    Delete       string   `json:"delete"`
    TermUpdate   string   `json:"termUpdate"`
    RateUpdate   string   `json:"rateUpdate"`
    Terminate    Terminate `json:"terminate"`
}

type DealStatus struct {
    Open      string   `json:"OPEN"`
    Delete    string   `json:"DELETE"`
    Close     string   `json:"CLOSE"`
    Terminate Terminate `json:"TERMINATE"`
    Error     string   `json:"ERROR"`
    Cancel    string   `json:"CANCEL"`
    WaitingToOpen string `json:"WAITING_TO_OPEN"`
}

type MidasType struct {
    ProductNumber   string `json:"productNumber"`
    MidasDealType   string `json:"midasDealType"`
    MidasDealSubType string `json:"midasDealSubType"`
}

type CoverageTransitAccount struct {
    Corporate string `json:"corporate"`
    Retail    string `json:"retail"`
}

type BaseRateList struct {
    RUB []BaseRate `json:"RUB"`
}

type BaseRate struct {
    BaseRateGroupCode string `json:"baseRateGroupCode"`
    BaseRateCode      string `json:"baseRateCode"`
}