package main

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type dataCharacteristics struct {
	AE         int64   `dynamodbav:"ae"`
	AM         int64   `dynamodbav:"-"`
	AMALF      uint64  `dynamodbav:"-"`
	AMALM      uint64  `dynamodbav:"-"`
	AS         int64   `dynamodbav:"-"`
	ASALF      uint64  `dynamodbav:"-"`
	ASALM      uint64  `dynamodbav:"-"`
	BL         int64   `dynamodbav:"-"`
	BLALF      uint64  `dynamodbav:"-"`
	BLALM      uint64  `dynamodbav:"-"`
	CHARTER    string  `dynamodbav:"charter"`
	FRELCH     int64   `dynamodbav:"-"`
	FTE        float64 `dynamodbav:"fte"`
	G01        int64   `dynamodbav:"g01"`
	G02        int64   `dynamodbav:"g02"`
	G03        int64   `dynamodbav:"g03"`
	G04        int64   `dynamodbav:"g04"`
	G05        int64   `dynamodbav:"g05"`
	G06        int64   `dynamodbav:"g06"`
	G07        int64   `dynamodbav:"g07"`
	G08        int64   `dynamodbav:"g08"`
	G09        int64   `dynamodbav:"g09"`
	G10        int64   `dynamodbav:"g10"`
	G11        int64   `dynamodbav:"g11"`
	G12        int64   `dynamodbav:"g12"`
	G13        int64   `dynamodbav:"g13"`
	GSHI       string  `dynamodbav:"gshi"`
	GSLO       string  `dynamodbav:"gslo"`
	HI         int64   `dynamodbav:"-"`
	HIALF      uint64  `dynamodbav:"-"`
	HIALM      uint64  `dynamodbav:"-"`
	HP         int64   `dynamodbav:"-"`
	HPALF      uint64  `dynamodbav:"-"`
	HPALM      uint64  `dynamodbav:"-"`
	ID         int64   `dynamodbav:"id"`
	KG         int64   `dynamodbav:"kg,omitempty"`
	LAT        float64 `dynamodbav:"lat"`
	LCITY      string  `dynamodbav:"lcity"`
	LATCOD     float64 `dynamodbav:"latcod,omitempty"`
	LEAID      int64   `dynamodbav:"leaid"`
	LEANAME    string  `dynamodbav:"leaName"`
	LONCOD     float64 `dynamodbav:"loncod,omitempty"`
	LONG       float64 `dynamodbav:"long,omitempty"`
	LSTATE     string  `dynamodbav:"lstate"`
	LSTREET1   string  `dynamodbav:"lstreet1"`
	LSTREET2   string  `dynamodbav:"lstreet2"`
	LZIP       int64   `dynamodbav:"lzip"`
	LZIP4      int64   `dynamodbav:"lzip4"`
	MAGNET     string  `dynamodbav:"magnet"`
	MEMBER     int64   `dynamodbav:"member"`
	NMCNTY     string  `dynamodbav:"nmcnty"`
	PHONE      string  `dynamodbav:"phone"`
	PREKST     int64   `dynamodbav:"prekst,omitempty"`
	REDLCH     int64   `dynamodbav:"-"`
	SCHLEVEL   string  `dynamodbav:"schLevel"`
	SCHNAME    string  `dynamodbav:"schName"`
	SCHTYPETXT string  `dynamodbav:"schTypeTxt"`
	STABR      string  `dynamodbav:"stabr"`
	STATUS     int64   `dynamodbav:"status"`
	STITLE1    string  `dynamodbav:"stitle1,omitempty"`
	STLEAID    string  `dynamodbav:"stleaid"`
	STUTERATIO float64 `dynamodbav:"stuteratio"`
	SYSTATUS   string  `dynamodbav:"sysStatus"`
	TITLE1     string  `dynamodbav:"title1,omitempty"`
	TOTFENROL  string  `dynamodbav:"-"`
	TOTFRL     int64   `dynamodbav:"-"`
	TOTMENROL  string  `dynamodbav:"-"`
	TOTALST    int64   `dynamodbav:"totalst"`
	TR         int64   `dynamodbav:"-"`
	TRALF      uint64  `dynamodbav:"-"`
	TRALM      uint64  `dynamodbav:"-"`
	UG         int64   `dynamodbav:"ug"`
	ULOCALE    string  `dynamodbav:"ulocale"`
	VIRTUAL    string  `dynamodbav:"virtual"`
	WH         int64   `dynamodbav:"-"`
	WHALM      uint64  `dynamodbav:"-"`
	WHALF      uint64  `dynamodbav:"-"`
	YEAR       string  `dynamodbav:"year"`
}

func getS3Object(cfg *aws.Config) *s3.GetObjectOutput {
	bucket := os.Getenv("AWS_MIGRATION_BUCKET")
	key := os.Getenv("AWS_MIGRATION_OBJECT")

	s3Client := s3.NewFromConfig(*cfg)

	data, e := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	if e != nil {
		log.Fatalf(e.Error())
	}

	return data
}

func connectDb(creds *aws.Config, tableName string) *dynamodb.Client {
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")

	identity, _ := creds.Credentials.Retrieve(context.TODO())

	cfg, e := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: endpoint}, nil
			})),

		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     identity.AccessKeyID,
				SecretAccessKey: identity.SecretAccessKey,
				SessionToken:    identity.SessionToken,
			},
		}))

	if e != nil {
		log.Fatalf("fatal error - %s", e.Error())
	}

	client := dynamodb.NewFromConfig(cfg)
	tableExists := hasTable(client, tableName)

	if tableExists == false {
		_ = createTable(client, tableName)
	}

	return client
}

func createTable(client *dynamodb.Client, tableName string) *dynamodb.CreateTableOutput {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("lzip"),
				AttributeType: types.ScalarAttributeTypeN,
			},
			{
				AttributeName: aws.String("schName"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("lzip"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("schName"),
				KeyType:       types.KeyTypeRange,
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	}

	out, e := client.CreateTable(context.TODO(), input)
	if e != nil {
		log.Fatalf("Error creating table %v", e)
	}

	return out
}

func hasTable(client *dynamodb.Client, tableName string) bool {
	var exists = true

	_, e := client.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	if e != nil {
		var notFoundEx *types.ResourceNotFoundException

		if errors.As(e, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", tableName)
		} else {
			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", tableName, e)
		}

		exists = false
	}

	return exists
}

func getCredentials() *aws.Config {
	cfg, e := config.LoadDefaultConfig(context.TODO())
	if e != nil {
		log.Fatalf(e.Error())
	}

	roleArn := os.Getenv("AWS_ROLE_TO_ASSUME")

	client := sts.NewFromConfig(cfg)
	creds := stscreds.NewAssumeRoleProvider(client, roleArn)
	cfg.Credentials = aws.NewCredentialsCache(creds)

	return &cfg
}

func closeConnection(seedData *s3.GetObjectOutput) {
	func(Body io.ReadCloser) {
		e := Body.Close()
		if e != nil {
			log.Fatalf(e.Error())
		}
	}(seedData.Body)
}

func writeItem(wid int, client *dynamodb.Client, processedItems <-chan dataCharacteristics, tableName string, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range processedItems {
		record, e := attributevalue.MarshalMap(v)
		log.Println("WriteItemWorker", wid, "started job", v, "with", len(processedItems), "jobs left")

		if e != nil {
			log.Printf("Error marshalling record %v", record)
		}

		out, e := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
			Item:      record,
			TableName: aws.String(tableName),
		})
		if e != nil {
			log.Printf("Error storing record %v", record)
		}

		log.Printf("Put item result %v", out)
	}
}

func main() {

	var wg sync.WaitGroup
	const dbItems = 100722

	extractCh := make(chan []string, dbItems)
	processCh := make(chan dataCharacteristics, dbItems)
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	cfg := getCredentials()
	seedData := getS3Object(cfg)
	defer closeConnection(seedData)

	dbClient := connectDb(cfg, tableName)

	reader := csv.NewReader(seedData.Body)
	parseHeader(reader)

	for workerId := 1; workerId <= 6; workerId++ {
		go processCSV(workerId, extractCh, processCh)
	}

	for i := 1; i <= dbItems; i++ {
		item, e := reader.Read()
		if e != nil {

			if numFields, ok := e.(*csv.ParseError); ok && numFields.Err == csv.ErrFieldCount {
				log.Printf("Item Number: %d - has incorrect number of fields", i)
				log.Fatalf(e.Error())
			} else {
				log.Fatalf(e.Error())
			}

		}

		extractCh <- item
	}

	close(extractCh)

	for workId := 1; workId <= 3; workId++ {
		wg.Add(1)
		go writeItem(workId, dbClient, processCh, tableName, &wg)
	}

	wg.Wait()
}

func parseHeader(h *csv.Reader) {
	_, e := h.Read()

	if e != nil {
		log.Fatalf(e.Error())
	}

	// return header
}

func processCSV(wid int, ch <-chan []string, processCh chan<- dataCharacteristics) {
	for item := range ch {
		log.Println("worker", wid, "started job", item[3], "with", len(ch), "jobs left")

		ae, _ := strconv.ParseInt(item[72], 0, 0)
		am, _ := strconv.ParseInt(item[40], 0, 0)
		amalf, _ := strconv.ParseUint(item[55], 0, 0)
		amalm, _ := strconv.ParseUint(item[54], 0, 0)
		as, _ := strconv.ParseInt(item[76], 0, 0)
		asalf, _ := strconv.ParseUint(item[57], 0, 0)
		asalm, _ := strconv.ParseUint(item[56], 0, 0)
		bl, _ := strconv.ParseInt(item[42], 0, 0)
		blalf, _ := strconv.ParseUint(item[61], 0, 0)
		blalm, _ := strconv.ParseUint(item[60], 0, 0)
		id, _ := strconv.ParseInt(item[3], 0, 0)
		frelch, _ := strconv.ParseInt(item[21], 0, 0)
		fte, _ := strconv.ParseFloat(item[46], 64)
		g01, _ := strconv.ParseInt(item[25], 0, 0)
		g02, _ := strconv.ParseInt(item[26], 0, 0)
		g03, _ := strconv.ParseInt(item[27], 0, 0)
		g04, _ := strconv.ParseInt(item[28], 0, 0)
		g05, _ := strconv.ParseInt(item[29], 0, 0)
		g06, _ := strconv.ParseInt(item[30], 0, 0)
		g07, _ := strconv.ParseInt(item[31], 0, 0)
		g08, _ := strconv.ParseInt(item[32], 0, 0)
		g09, _ := strconv.ParseInt(item[33], 0, 0)
		g10, _ := strconv.ParseInt(item[34], 0, 0)
		g11, _ := strconv.ParseInt(item[35], 0, 0)
		g12, _ := strconv.ParseInt(item[36], 0, 0)
		g13, _ := strconv.ParseInt(item[37], 0, 0)
		hi, _ := strconv.ParseInt(item[41], 0, 0)
		hialf, _ := strconv.ParseUint(item[59], 0, 0)
		hialm, _ := strconv.ParseUint(item[58], 0, 0)
		hp, _ := strconv.ParseInt(item[44], 0, 0)
		hpalf, _ := strconv.ParseUint(item[65], 0, 0)
		hpalm, _ := strconv.ParseUint(item[64], 0, 0)
		kg, _ := strconv.ParseInt(item[24], 0, 0)
		lat, _ := strconv.ParseFloat(item[0], 64)
		latcod, _ := strconv.ParseFloat(item[47], 64)
		leaid, _ := strconv.ParseInt(item[6], 0, 0)
		loncod, _ := strconv.ParseFloat(item[48], 64)
		long, _ := strconv.ParseFloat(item[1], 64)
		lzip, _ := strconv.ParseInt(item[14], 0, 0)
		lzip4, _ := strconv.ParseInt(item[15], 0, 0)
		member, _ := strconv.ParseInt(item[39], 0, 0)
		prekst, _ := strconv.ParseInt(item[23], 0, 0)
		redlch, _ := strconv.ParseInt(item[22], 0, 0)
		status, _ := strconv.ParseInt(item[70], 0, 0)
		stuteratio, _ := strconv.ParseFloat(item[51], 64)
		totalst, _ := strconv.ParseInt(item[38], 0, 0)
		totfrl, _ := strconv.ParseInt(item[20], 0, 0)
		tralf, _ := strconv.ParseUint(item[67], 0, 0)
		tralm, _ := strconv.ParseUint(item[66], 0, 0)
		tr, _ := strconv.ParseInt(item[45], 0, 0)
		ug, _ := strconv.ParseInt(item[71], 0, 0)
		wh, _ := strconv.ParseInt(item[43], 0, 0)
		whalf, _ := strconv.ParseUint(item[63], 0, 0)
		whalm, _ := strconv.ParseUint(item[62], 0, 0)

		record := dataCharacteristics{
			AE:         ae,
			AM:         am,
			AMALF:      amalf,
			AMALM:      amalm,
			AS:         as,
			ASALF:      asalf,
			ASALM:      asalm,
			BL:         bl,
			BLALF:      blalf,
			BLALM:      blalm,
			CHARTER:    item[77],
			FRELCH:     frelch,
			FTE:        fte,
			G01:        g01,
			G02:        g02,
			G03:        g03,
			G04:        g04,
			G05:        g05,
			G06:        g06,
			G07:        g07,
			G08:        g08,
			G09:        g09,
			G10:        g10,
			G11:        g11,
			G12:        g12,
			G13:        g13,
			GSHI:       item[18],
			GSLO:       item[17],
			HI:         hi,
			HIALF:      hialf,
			HIALM:      hialm,
			HP:         hp,
			HPALF:      hpalf,
			HPALM:      hpalm,
			ID:         id,
			KG:         kg,
			LAT:        lat,
			LCITY:      item[12],
			LATCOD:     latcod,
			LEAID:      leaid,
			LEANAME:    item[8],
			LONCOD:     loncod,
			LONG:       long,
			LSTATE:     item[13],
			LSTREET1:   item[10],
			LSTREET2:   item[11],
			LZIP:       lzip,
			LZIP4:      lzip4,
			MAGNET:     item[78],
			MEMBER:     member,
			NMCNTY:     item[50],
			PHONE:      item[16],
			PREKST:     prekst,
			REDLCH:     redlch,
			SCHLEVEL:   item[75],
			SCHNAME:    item[9],
			SCHTYPETXT: item[73],
			STABR:      item[5],
			STATUS:     status,
			STITLE1:    item[53],
			STLEAID:    item[7],
			STUTERATIO: stuteratio,
			SYSTATUS:   item[74],
			TITLE1:     item[52],
			TOTFENROL:  item[69],
			TOTFRL:     totfrl,
			TOTMENROL:  item[68],
			TOTALST:    totalst,
			TR:         tr,
			TRALF:      tralf,
			TRALM:      tralm,
			UG:         ug,
			ULOCALE:    item[49],
			VIRTUAL:    item[19],
			WH:         wh,
			WHALM:      whalm,
			WHALF:      whalf,
			YEAR:       item[4],
		}
		log.Println("worker", wid, "     finished job", record.ID)
		log.Println(record.ID, record.LSTATE, record.NMCNTY, record.LZIP, record.SCHNAME)
		processCh <- record
	}
}
