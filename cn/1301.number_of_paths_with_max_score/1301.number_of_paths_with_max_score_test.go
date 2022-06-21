package cn

import (
	"reflect"
	"testing"
)

func Test_pathsWithMaxScore(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{board: []string{"E23", "2X2", "12S"}}, []int{7, 1}},
		{"", args{board: []string{"E12", "1X1", "21S"}}, []int{4, 2}},
		{"", args{board: []string{"E11", "XXX", "11S"}}, []int{0, 0}},
		{"", args{board: []string{"EX", "XS"}}, []int{0, 1}},
		{"", args{board: []string{"E4789338X943596124X2676X552X587877X456943458X29735", "611684759486631913932337237231351921X2152919376427", "4499519117827344997451XX34X46693XX7181343557483669", "414951X685152X89829782685X4912581351X3216914721551", "X387271851925X3629265X99195X5897581179XX369637813X", "1X8X2682518937289551X98X7983XX34993116413343558825", "X92X12119593186X675113X682143777XX8981619298251984", "X671798198463X5314971262X9392393XXX544537813812728", "81856146535454X3678775784456289257XX8221X2488XXX68", "77X3592XX94844399282X2X6336122XX7X18244862821X26XX", "28885X948512X3585X27824186222X73X9X56441X9X4689517", "344X495X682875968X82X9877379XX386748175X6293X44159", "1924352186149295919715X27X555626X17798524189528625", "3435681879X49727366745492X648X5952772978787143263X", "412933788234154913356X2X9X144818X21XX5629259785133", "489644765X456XX44XX2X5387637879X662941398337817381", "7617826679176XX173173537173164967296764519X3427693", "7X69X5466277665871135253486758156766536X5436X16728", "318152574426X696X18X833396113X31862234511611X89691", "X147492241256344555237X94772X9X5136226469551942X29", "X29846X49419853778154XX636X35XX5232391787617416258", "8885851X996538X163323347235993741926591X72X1761X14", "226X8136988863232963682217521777419144333838517835", "3197757518X241117949451X423XX55861XX6161938551X752", "49X149355329X617X51X21965452X962X42762X25968X73754", "X1176483834957794897816132X5X942366794665183797399", "48294674492176X1X663644X58X17729X4482638X92X482422", "51X16X157889846864X1436338776687677X44895154219626", "782XX82XX432848434721692X55564975938X2649681569663", "26792517214X8863X9896XX64619817916123168945761X526", "X7XX319393797X184679X421943743279X1486126364341595", "11425763X68563511X2496983325426X151456X5X459542989", "89872654932254X5525587692326476X65562X59X635129314", "6X944371656471737X9673645X4145X821X942995885X86522", "64XX328X8243XX3445X6412955X87X42355234X73223243421", "944391299X158X962X9X259552884441434XX9349X5XX79855", "98422711429356771336494176X56X2584376X2354X72X5416", "7135X7X9X3X699929714X61664916968X1896X5X1985386642", "22969X77414X2154167176388X3313X918X1558161XX413862", "4X141958614412616921588565488847635X837996835937X6", "X416X64649X955369739187781488XX77129X6966899351X74", "9X2349175X345X7469265842X591X4748167996963X63X7211", "3117349374592412365636726147X52469X8921X1888159627", "254513X73629359514989286822X7X89391797X357546X9145", "8882877128738295914941X56995X243888XX595873XX99217", "X947591X382X19613453263544415821689719794546655278", "51X4565271954965486X492693X731844471486X149X7166X5", "731X1816268181645946X2123376981X13X834338226X28X36", "6X72X3599546159X378191718821X746789239488712584143", "5982153336X61729X66339596838X77751396645929982913S"}}, []int{649, 12}},
		{"", args{board: []string{"E81X9983419X57565746547815665454989524693X2X97X495745461552X6459X15166812832325X21884262X32219286971", "1488266XX521X834427522XX87X128473498345884746X484969X97621781231166X518566468X113X6X46X6199X51112645", "878X172629248263633145X4931582381286963788584426927434937X447556692859351575X9565266144849X98527X315", "31936XX35282979X69887326419151518X17X4861246729X91X774X332754X3416497564457398X366162X39X9X19877X95X", "13275865544788X358X5X2349136X5927X762949634497779377994X98XXX451167613528X7176277275715X573744X934X4", "6383562X4X2X1531457331859994X5678141XX3262X925787947285549321411129X57588342993647321233187145X33371", "X17494325935691X27453729X7881X622221386421455426258X358647554X233134498732357X564773883X995938457X33", "64336241888836342221XXX68582518X548256124985958297654267X675X671X7699218891XX48955484393377378873553", "743918X225426498391717447711852189657529921137773957788863259358452X165538791X8383459688X99351636433", "22988156139392482153466XX614X565XX493418838X62X9X492X36814284X453987X186466X44896X85944476X4959772X8", "X33142565759X156127848317286X646413X6234X697X84213X8181797785X83735655786497893416748481X2543XX89286", "977973417X4227394665977618464597976X842673X9319167X94131916879115579447717X8235488758376897157388612", "895965X544615X531X871284915816769X19637X83173834661421655X2384XX938746X41632177142197915X456X91X7X36", "12262X368323459498266X29426279144X565513671X39328891597785674749942537764399626373459X3889896773X599", "6771XX76973X1583819X775724423927X892X856X4X2744785X7643739311XX93896199X4X27X984X41768482X8519247879", "4964425393168817X494X23X145611X86X9X75678225756496463161664X7878897468842999777966697479497987459476", "9149337794575767327244858691569589398346531944125955X39187566917498532119X2367X8133746927X28916XX76X", "887881239173852572659X54X7185X78549237126215856697116394587389732524831XX2676975X659249X438682669277", "25238X6X374466228781856265767321225X3875968418497759X2683X9631299844583378X2658977499X37695567927868", "465195347X4488417645796683398345343636741849661967X461247785212X489933735331888777755798867967614315", "4215738836475864175912956613551718339738442X421389946896645X72118X7615245X357815482911694859427329XX", "2X5975418486X25X65X963172841775953219783638X25X676X38848269X9428839492297578244481719389454395611466", "93374829X959X5261788988847544X7353X49698994741863692999X788716278314X96313652741139X48235139333X9X46", "36961XX81931887X769984885717X62783533552245193577379867113X363X943X721377X62536893265X7338824X5X6133", "X815642X9278481X14X527923391692646667158876876698417629297292527X842321345X45467511X5XX6287363613594", "195X499414465X5631885148358794597X4928113881781674746X2714864238157866681151X248186X556674X513219317", "36X3561X6733188X911565X31823X943X9111366232462958211128XX162668764123X471X125X739384113926685XX28X82", "399674113128757438933327519863534648437971133X532244231351722X9665X3322571253993252215X1839684797879", "7619996X7789X51X9541X3968X4X414X13XX45X5X798X625367X5X9267839X2461298187X9X24X4189446559282X437367X6", "328959292172321624295954498441XX7XX83X941849X86958931439133X4866123282X315822765946263175454X2X25137", "9X9444998691799658278696726475666X849X5378246791874415734X273818311X484X2148214477933X44789394494552", "952859314496739885X53343848323X531X6493X9X88959558979868279925X57X37X57331171277513X25XX684174294674", "38X94X3746673151957311557X4331369813X833X99X596863X8X29X7X7322526593837777X6X4X11227188385X932851977", "97184X6841686735284X492528X176X25839X695855X14465597656X3817826656263715686X956742761659X66X61359868", "3X18991X2X3259451964273X23X4832622368585284242722X1732677X31152132929X4454712372254822861114X6498379", "48498851486X22178561417XX6115511456X8282565878484834774781643X1913485242556456382618337828228258331X", "X59631982X687776167426224XX149X6877214364825499756375XX7512968954957X716339191X35XX29668X166438X9X41", "2743XX4252595635929872453615X82684326X4X394X7169596X21114164885233636X95472X8147965754X4512X2159X8XX", "61477115876X1X39166146XXX882X8X958X75612X769663358167688X71497X75X1X8933753278827X1392XX23961X592655", "92278X2999758X12X433XX789X16X4XX231315982384915728X742125729188137868X26X12189X4393633653X7367671236", "884545XXX863379X75928963852X176622X6254889X82X483147652656795X946X856924353727342973463871726X31X652", "8X37349727X353188XX2144X893723436694X4132927X183281626X2X878X7899766418626X751944244786423128X5X89X2", "16341668541735XX3984546217694958134471416157X7138X14X481387X4794861332175495266X88688339114396145992", "536681436154546613763X36237741X6449553X941699734693929672265X846X1536291X6191X37X5236265614944259141", "319X7811171256354X8995X971992699X677919X4X77825779715194862X99466146852X8X6X98874529898XX94973161675", "3113998X6X26XX1877891263X2X9226245281785288729X83827X4615741481428331855661X524322939465946XX8645767", "76784575155X97865X38245X39127776844249X3348683XX599728774914124975771426799X83498715X99799116XXX5X58", "897625794127756X6442859698429318196882X9663916943752286544X2773189837722227944X588837149462625998753", "313X632158X6X719281X77125848776971951531552X6778X58636X38344X49843489441592456633X9717352818437431X3", "345713175899966782781754373X54X24X52X2X6585785155979727861516X341637X376X9X48169X8827315479794828261", "X77X81686X9X4224X99859763994354X57856259618848498151X34646X643278154176431775716526771465X5553631924", "6972985996X59XX2X536X983X2223996972X51X18X73786491X256253371573999X62637X874346149X83X42245371559465", "71782542248283X17712424474289425X914818663656X424717141214897151X6765562748X353717846236323622197X29", "454218828XX16931374366287361115385X876261119353875959725X2291184488684792X96435X124XX697X248X6766878", "727X6X786184788669236X3X779X824X1182685952X2414684X4283577865X29745X46486354887189X71798X995238349X2", "95272735X9636579166578X48724386611385XX584X2378X9X167862X8645X1X92X611425XX55X95784224799641923X5519", "31698X99156268629431142459832556X828315422475X2815597897672334172X5598386836X9864X819X6XX85294525896", "2341X51411X333645234179375766148447X1691613556X692897313X2291427369887145384678666926389945917221X3X", "5567763X9759541462365141486524151285613547484415583X774812595X2X8733787919443633538859897X32915414X6", "525161X753892926833122253515144653287566965511X95X5917853981443X548278627X21269931392751557246331416", "2659271637111X845151179356294532X2781717636887X7583911198478X371346479648XX31483227791787X287X519144", "X413233936351X811549243535X5273X747X6X397262692X6432369854X5716X351822795938719479716518X358147343X3", "71961751582496871467362482224974278162157339925598888428914X893586694X16X542X33479XXX4397161481X4474", "338774X2X8X7915917356681X7712XXX336993324112445434791X11X33145X91714273943X742327X232923926211228718", "1786X521X111X8X8X3574113422X48459538937X29879811887299443533X2X36334735337XX5X6828268X61368984553X65", "9434X28X66X912X216X62715353337186825152X21213339257325293327397X416544272247384793935748788922324628", "15358637135988842188X141X7614215452919X225X86858X9642463561X61153662934924379149948141X4X426381XX3X6", "72785643X3633X969477789XX6XX71532X56553913524X129198XX4774279217689297359574X98275488145363891397474", "2914741571854X7469X7X645225919653522X44165X325253784187459X8536755478X66X79752484618X2882956495X89X6", "168431241X59881439226366716X647796262X237374712X599775327388X9188255791996768752762623122X32371X4498", "158486787285X7827138417433775619911X38797538562X474596XX51497762818926197354827747634329513X32727751", "484887199X7269653145X62764724937657545644X91X24X659924815X97969748121X48X997X44611X3364491X7363X2436", "693581838X54784935381X1X67271X3467245X5888559X53949232748833719822X994849376156X311493X1912358847796", "1627919376775973536355997475XX5928317417259724649861X287477552819X493375783X269X2X821256573252429X55", "82891813254862X32776863448772859569639388662455369673369364231284941628X73529871423XX476549X322XX168", "4XX113185X5714X943X5X996736X48476699667527X58539889359295854X246939975482383X89XXX8X95157253833162X6", "3685558X2618X4662467251145X1165791395667343699737916925939479X7X6525XX96526X15195434671878X417965713", "68922X1222X76X6244X527212X49318179X4X537576759941437369579782252XXX1242873912154X1552679587889196X61", "1X55696114674929634939158X4X751X4283X97259X999414X58142228691595487284327896X4X94X372541763864X25859", "243687243565363698X1X65767766211X1589122323118X219282X62X9892397785535982X3815283779711693X1X1514585", "12XX9X72X235X243252X913197891X673456X37135657535938X4X67267898746894545656827X371399323195718536X865", "3613381936726223X4549565812XX93988114166736265289X6X66284473589X26553881355927316X4178464337X7792816", "X5116915355832452884585691672773561622318276333328441141819811197934843825642954469553238X8744166763", "3491468365X52X265197739X8227993784835443993X788X26218667263244X93492X2682897893397X698X2X59257691595", "525672189X4385684391171768916354528135477X7828597XX572X3798X7X36484659594311734773116499466746267549", "4533679X587516625522751X146251229X34X1142417967X6285X1925X5712542156X8X2X7398X51X23X263XX62795X43748", "2178415839641787232172289X564X758X63437616111268574X83455225572869754878X91X4159X281398935826X35XX58", "5X846749378168XXX7396369941731348471191277985524X3584784917674622944188X7364687583X7923X66XX271X638X", "63971897X461965789XX893XX12182125987832767489XX7X9393352711X37567966X664316921226431X74233X327154765", "896314682X382354X388X33949975X81814671X4X215X68126597975352328X23883549247X897567161358X5X413X431245", "84453X22456676524926X2561X779948126226412398X6497419598863913774395688442689484169676321X387488378X5", "769847217943513847315476386215X9827877613144317X647916348941998X95373X68X29443X364627469174X5X387676", "6875X8255562527878X714637486796641477373127122X41896578X7132X328879628X889995137474848X2X27217XX7384", "X2398X41728387326699749X71985873262622953725X8X3X838196747248X319X21512788914418161359767796942X92X2", "79X1594X879X78877X8657754711255283655186X7743862456X94291X4X81126257185665X7938458113244975775118881", "968814792658X3X457359644684581977352487334289411464X552324334X4938X1511145434X224749172X971925356258", "737258X63X8316861245363786672X814132X5141967976981535X6795624177437X64639179448215616X23555638419865", "257582322863519491454925882137897651X87698X57217858712472235979X62441343696912162219145X134X3569X479", "348118X446494X44144294X947X2197887324X71141832XX29793433X65977224X44781378X7226X654471X3373853365881", "X872659927767353X8276671383973816782338543X8231259837X37651X8723448267836846813265145X2682X544X7X8XS"}}, []int{1352, 130}},
		{"", args{board: []string{"E999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999S"}}, []int{1773, 690285631}},
		{"", args{board: []string{"E111", "1111", "1111", "111S"}}, []int{5, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathsWithMaxScore(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathsWithMaxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
