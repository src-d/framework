package ecosystem

import "srcd.works/framework/utils"

type Ecosystems map[string]float64

func (e Ecosystems) Add(eco string, size float64) {
	e[eco] = e[eco] + size
}

func (e Ecosystems) SuitableLanguages() map[string]map[string]bool {
	o := make(map[string]map[string]bool, 0)

	for k, _ := range e {
		l, ok := ecosystemLanguages[k]
		if !ok {
			continue
		}

		for _, lang := range l {
			if _, ok := o[lang]; !ok {
				o[lang] = make(map[string]bool, 0)
			}

			o[lang][k] = true
		}
	}

	return o
}

var ecosystemLanguages = map[string][]string{
	Unity: []string{"C#"},

	Gtk:       []string{"C", "C++", "C#", "Go", "Java", "Python", "Ruby"},
	Qt:        []string{"C++", "C#", "Go", "Java", "Python", "Ruby"},
	WxWidgets: []string{"C++", "Go", "Java", "PHP", "Python", "Ruby"},

	Catch:      []string{"C++"},
	GoogleTest: []string{"C++"},

	Ginkgo:    []string{"Go"},
	Gocheck:   []string{"Go"},
	Goconvey:  []string{"Go"},
	GoTestify: []string{"Go"},

	Grails:         []string{"Groovy"},
	GroovyUnitTest: []string{"Groovy"},

	QuickCheck: []string{"Haskell"},

	Android:        []string{"Java"},
	Deeplearning4j: []string{"Java"},
	Encog:          []string{"Java"},
	GATE:           []string{"Java"},
	GWT:            []string{"Java"},
	JavaML:         []string{"Java"},
	JMonkeyEngine:  []string{"Java"},
	JSAT:           []string{"Java"},
	JSFML:          []string{"Java"},
	JUnit:          []string{"Java"},
	LibGdx:         []string{"Java"},
	LingPipe:       []string{"Java"},
	LWJGL:          []string{"Java"},
	MALLET:         []string{"Java"},
	Mockito:        []string{"Java"},
	OpenNLP:        []string{"Java"},
	PlayFramework:  []string{"Java"},
	Slick2D:        []string{"Java"},
	SpringMVC:      []string{"Java"},
	Struts:         []string{"Java"},
	TestNG:         []string{"Java"},
	Vaadin:         []string{"Java"},
	Wicket:         []string{"Java"},

	AngularJS:        []string{"JavaScript", "CoffeeScript", "TypeScript"},
	BabylonJS:        []string{"JavaScript", "CoffeeScript", "TypeScript"},
	BackboneJS:       []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Brick:            []string{"JavaScript", "CoffeeScript", "TypeScript"},
	BusterJS:         []string{"JavaScript", "CoffeeScript", "TypeScript"},
	ChaplinJS:        []string{"JavaScript", "CoffeeScript", "TypeScript"},
	ChromeExtension:  []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Cocos2DHTML5:     []string{"JavaScript", "CoffeeScript", "TypeScript"},
	CraftyJS:         []string{"JavaScript", "CoffeeScript", "TypeScript"},
	D3:               []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Dojo:             []string{"JavaScript", "CoffeeScript", "TypeScript"},
	EmberJS:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	ExpressJS:        []string{"JavaScript", "CoffeeScript", "TypeScript"},
	FirefoxExtension: []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Intern:           []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Jasmine:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Jest:             []string{"JavaScript", "CoffeeScript", "TypeScript"},
	JQuery:           []string{"JavaScript", "CoffeeScript", "TypeScript"},
	KiwiJS:           []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Knockout:         []string{"JavaScript", "CoffeeScript", "TypeScript"},
	KoaJS:            []string{"JavaScript", "CoffeeScript", "TypeScript"},
	LocomotiveJS:     []string{"JavaScript", "CoffeeScript", "TypeScript"},
	MatreshkaJS:      []string{"JavaScript", "CoffeeScript", "TypeScript"},
	MelonJS:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	MochaJS:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	MoTools:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	NodeJS:           []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Phaser:           []string{"JavaScript", "CoffeeScript", "TypeScript"},
	PixiJS:           []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Polymer:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Prototype:        []string{"JavaScript", "CoffeeScript", "TypeScript"},
	ReactJS:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	SailsJS:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	SenecaJS:         []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Sinon:            []string{"JavaScript", "CoffeeScript", "TypeScript"},
	SocketIO:         []string{"JavaScript", "CoffeeScript", "TypeScript"},
	Spine:            []string{"JavaScript", "CoffeeScript", "TypeScript"},
	StageJS:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	TotalJS:          []string{"JavaScript", "CoffeeScript", "TypeScript"},
	UnderscoreJS:     []string{"JavaScript", "CoffeeScript", "TypeScript"},
	UnitJS:           []string{"JavaScript", "CoffeeScript", "TypeScript"},
	VelocityJS:       []string{"JavaScript", "CoffeeScript", "TypeScript"},
	YUI:              []string{"JavaScript", "CoffeeScript", "TypeScript"},
	MeteorJS:         []string{"JavaScript", "CoffeeScript", "TypeScript"},

	Aura:         []string{"PHP"},
	CakePHP:      []string{"PHP"},
	Codeception:  []string{"PHP"},
	CodeIgniter:  []string{"PHP"},
	Drupal7:      []string{"PHP"},
	Joomla:       []string{"PHP"},
	Kohana:       []string{"PHP"},
	Laravel:      []string{"PHP"},
	Phalcon:      []string{"PHP"},
	PHPSpec:      []string{"PHP"},
	PHPUnit:      []string{"PHP"},
	Silex:        []string{"PHP"},
	Slim:         []string{"PHP"},
	Symfony:      []string{"PHP"},
	Typo3:        []string{"PHP"},
	Wordpress:    []string{"PHP"},
	Yii2:         []string{"PHP"},
	Yii:          []string{"PHP"},
	Zend2:        []string{"PHP"},
	Zend:         []string{"PHP"},
	PHPExtension: []string{"C"},

	Bluebream:      []string{"Python"},
	Bottle:         []string{"Python"},
	Caffe:          []string{"Python"},
	Celery:         []string{"Python"},
	CherryPy:       []string{"Python"},
	Cocos2dPython:  []string{"Python"},
	Django:         []string{"Python"},
	Flask:          []string{"Python"},
	Grok:           []string{"Python"},
	Keras:          []string{"Python"},
	Mlpy:           []string{"Python"},
	NLTK:           []string{"Python"},
	Numpy:          []string{"Python"},
	Pandas:         []string{"Python"},
	PyBrain:        []string{"Python"},
	Pygame:         []string{"Python"},
	Pyglet:         []string{"Python"},
	PyML:           []string{"Python"},
	PyMunk:         []string{"Python"},
	Pyramid:        []string{"Python"},
	PyTest:         []string{"Python"},
	PythonGuava:    []string{"Python"},
	PythonSFML:     []string{"Python"},
	PythonUnittest: []string{"Python"},
	ScikitLearn:    []string{"Python"},
	Scrapy:         []string{"Python"},
	TornadoWeb:     []string{"Python"},
	TurboGears:     []string{"Python"},
	Twisted:        []string{"Python"},
	Web2Py:         []string{"Python"},
	WebPy:          []string{"Python"},
	Zope:           []string{"Python"},

	Ai4r:        []string{"Ruby"},
	Bacon:       []string{"Ruby"},
	Camping:     []string{"Ruby"},
	Capybara:    []string{"Ruby"},
	Chingu:      []string{"Ruby"},
	Cuba:        []string{"Ruby"},
	Cutest:      []string{"Ruby"},
	Gosu:        []string{"Ruby"},
	GoTesting:   []string{"Ruby"},
	Minitest:    []string{"Ruby"},
	Padrino:     []string{"Ruby"},
	Pakyow:      []string{"Ruby"},
	RbLibsvm:    []string{"Ruby"},
	RSpec:       []string{"Ruby"},
	RubyBand:    []string{"Ruby"},
	Rubygame:    []string{"Ruby"},
	RubyOnRails: []string{"Ruby"},
	RubyOpenGL:  []string{"Ruby"},
	RubySDL:     []string{"Ruby"},
	Sinatra:     []string{"Ruby"},
	Spinach:     []string{"Ruby"},
	TestUnit:    []string{"Ruby"},
	Volt:        []string{"Ruby"},
	Chef:        []string{"Ruby"},

	ScalaTest: []string{"Scala"},

	Quick:     []string{"Swift"},
	XCodeTest: []string{"Swift", "Objective-C"},

	Tensorflow: []string{"Python", "C++"},

	MongoDb:    []string{"C", "C++", "PHP", "Java", "Python", "Ruby", "JavaScript", "CoffeeScript", "TypeScript", "Go"},
	MySQL:      []string{"C", "C++", "PHP", "Java", "Python", "Ruby", "JavaScript", "CoffeeScript", "TypeScript", "Go"},
	PostgreSQL: []string{"C", "C++", "PHP", "Java", "Python", "Ruby", "JavaScript", "CoffeeScript", "TypeScript", "Go"},
	SQLite:     []string{"C", "C++", "PHP", "Java", "Python", "Ruby", "JavaScript", "CoffeeScript", "TypeScript", "Go"},
	Redis:      []string{"C", "C++", "PHP", "Java", "Python", "Ruby", "JavaScript", "CoffeeScript", "TypeScript", "Go"},
	Cassandra:  []string{"C", "C++", "PHP", "Java", "Python", "Ruby", "JavaScript", "CoffeeScript", "TypeScript", "Go"},
	Neo4j:      []string{"C", "C++", "PHP", "Java", "Python", "Ruby", "JavaScript", "CoffeeScript", "TypeScript", "Go"},
	OracleDb:   []string{"C", "C++", "PHP", "Java", "Python", "Ruby", "JavaScript", "CoffeeScript", "TypeScript", "Go"},

	Docker: []string{"C", "C++", "C#", "Haskell", "PHP", "Java", "Scala", "Groovy", "Clojure", "Python", "Ruby", "JavaScript", "CoffeeScript", "TypeScript", "Go"},

	OpenCL:   []string{"C", "C++", "Java", "Python", "Ruby", "Javascript"},
	RabbitMQ: []string{"C++", "PHP", "Java", "Python", "Ruby", "Javascript", "Go"},
	Hadoop:   []string{"C++", "Ruby", "Go", "Python", "Java"},
	Spark:    []string{"Python", "Java", "Scala", "Ruby", "JavaScript", "CoffeeScript", "TypeScript"},

	Ansible: []string{"YAML"},

	AmazonAWS:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonEC2:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonCloudTrail: []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonGlacier:    []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonIAM:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonKMS:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonRDS:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonRoute53:    []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonS3:         []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonSNS:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	AmazonSQS:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},

	GoogleCloudPlatform:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudBigQuery:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudCompute:         []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudDatastore:       []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudDNS:             []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudErrorReporting:  []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudLanguage:        []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudLogging:         []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudMonitoring:      []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudPubSub:          []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudResourceManager: []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudSpeech:          []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudStorage:         []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudTrace:           []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudTranslate:       []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
	GoogleCloudVision:          []string{"C", "C++", "C#", "Clojure", "Go", "Groovy", "Java", "PHP", "Python", "Ruby", "Scala", "JavaScript", "CoffeeScript", "TypeScript"},
}

var languageEcosystems map[string][]string

func GetEcosystemLanguages(eco string) []string {
	return ecosystemLanguages[eco]
}

func GetLanguageEcosystems(lang string) []string {
	return languageEcosystems[lang]
}

func init() {
	languageEcosystems = utils.ReverseStringListMap(ecosystemLanguages)
}

const (
	Android          = "android"
	AngularJS        = "angular-js"
	ReactJS          = "react-js"
	D3               = "d3"
	Aura             = "aura"
	BackboneJS       = "backbone-js"
	CakePHP          = "cakephp"
	CodeIgniter      = "codeigniter"
	Dojo             = "dojo"
	Docker           = "docker"
	EmberJS          = "emberjs"
	ExpressJS        = "expressjs"
	JQuery           = "jquery"
	Knockout         = "knockout"
	KoaJS            = "koajs"
	Kohana           = "kohana"
	Laravel          = "laravel"
	LocomotiveJS     = "locomotive-js"
	MatreshkaJS      = "matreshka-js"
	NLTK             = "nltk"
	Keras            = "keras"
	Caffe            = "caffe"
	Pandas           = "pandas"
	Phalcon          = "phalcon"
	SailsJS          = "sailsjs"
	ScikitLearn      = "scikit-learn"
	Silex            = "silex"
	Slim             = "slim"
	SocketIO         = "socket-io"
	Spine            = "spine"
	Symfony          = "symfony"
	TotalJS          = "total-js"
	UnderscoreJS     = "underscore-js"
	YUI              = "yui"
	Yii              = "yii"
	Yii2             = "yii-2"
	Zend             = "zend"
	Zend2            = "zend-2"
	Drupal7          = "drupal-7"
	PHPExtension     = "php-extension"
	Polymer          = "polymer"
	Brick            = "brick"
	MoTools          = "motools"
	NodeJS           = "node-js"
	ChromeExtension  = "chrome-extension"
	FirefoxExtension = "firefox-extension"
	Unity            = "unity"
	Django           = "django"
	Flask            = "flask"
	Bottle           = "bottle"
	Pyramid          = "pyramid"
	Web2Py           = "web2py"
	WebPy            = "webpy"
	TurboGears       = "turbo-gears"
	CherryPy         = "cherry-py"
	Zope             = "zope"
	Grok             = "grok"
	Bluebream        = "bluebream"
	PythonGuava      = "guava"
	TornadoWeb       = "tornado-web"
	Twisted          = "twisted"
	Scrapy           = "scrapy"
	PythonUnittest   = "python-unittest"
	PyTest           = "pytest"
	SpringMVC        = "spring-mvc"
	GWT              = "gwt"
	JSF              = "jsf"
	Struts           = "struts"
	PlayFramework    = "play-framework"
	Wicket           = "wicket"
	Vaadin           = "vaadin"
	JavaML           = "java-ml"
	JSAT             = "jsat"
	OpenNLP          = "open-nlp"
	LingPipe         = "ling-pipe"
	GATE             = "gate"
	MALLET           = "mallet"
	Encog            = "encog"
	Deeplearning4j   = "deeplearning4j"
	JUnit            = "junit"
	TestNG           = "testng"
	JTiger           = "jtiger"
	Mockito          = "mockito"
	LWJGL            = "lwjgl"
	Slick2D          = "slick2d"
	LibGdx           = "libgdx"
	JMonkeyEngine    = "jmonkey-engine"
	JSFML            = "jsfml"
	Grails           = "grails"
	PyBrain          = "pybrain"
	PyML             = "pyml"
	Mlpy             = "mlpy"
	Pygame           = "pygame"
	Pyglet           = "pyglet"
	Cocos2dPython    = "cocos2dpython"
	PyMunk           = "pymunk"
	PythonSFML       = "python-sfml"
	PHPUnit          = "phpunit"
	PHPSpec          = "phpspec"
	Codeception      = "codeception"
	Jasmine          = "jasmine"
	MochaJS          = "mochajs"
	UnitJS           = "unitjs"
	BusterJS         = "busterjs"
	Sinon            = "sinon"
	Intern           = "intern"
	Jest             = "jest"
	VelocityJS       = "velocityjs"
	ChaplinJS        = "chaplinjs"
	Prototype        = "prototype"
	Enyo             = "enyo"
	BabylonJS        = "babylonjs"
	CraftyJS         = "craftyjs"
	Phaser           = "phaser"
	PixiJS           = "pixijs"
	PhysicsJS        = "physicsjs"
	Cocos2DHTML5     = "cocos2dhtml5"
	MelonJS          = "melonjs"
	StageJS          = "stagejs"
	KiwiJS           = "kiwijs"
	RubyOnRails      = "rubyonrails"
	Sinatra          = "sinatra"
	Volt             = "volt"
	Padrino          = "padrino"
	Pakyow           = "pakyow"
	Cuba             = "cuba"
	Camping          = "camping"
	RubyMotion       = "rubymotion"
	MobiRuby         = "mobiruby"
	RbLibsvm         = "rblibsvm"
	RubyBand         = "rubyband"
	Ai4r             = "ai4r"
	RSpec            = "rspec"
	Capybara         = "capybara"
	Cutest           = "cutest"
	Minitest         = "minitest"
	Bacon            = "bacon"
	TestUnit         = "testunit"
	Spinach          = "spinach"
	Gosu             = "gosu"
	RubyOpenGL       = "rubyopengl"
	Chingu           = "chingu"
	Rubygame         = "rubygame"
	RUDL             = "rudl"
	RubySDL          = "rubysdl"
	RubySFML         = "sfml"
	GoTesting        = "go-testing"
	Gocheck          = "gocheck"
	Ginkgo           = "ginkgo"
	GoTestify        = "go-testify"
	Goconvey         = "goconvey"
	XCodeTest        = "xcode-test"
	Quick            = "swift-quick"
	GroovyUnitTest   = "groovy-unittest"
	ScalaTest        = "scala-test"
	GoogleTest       = "google-test"
	Catch            = "catch"
	QuickCheck       = "quickcheck"
	ElasticSearch    = "elasticsearch"
	Lucene           = "lucene"
	Solr             = "solr"
	Sphinx           = "sphinx"
	Tensorflow       = "tensorflow"
	SenecaJS         = "senecajs"
	OpenCL           = "opencl"
	Hadoop           = "hadoop"
	Celery           = "celery"
	Joomla           = "joomla"
	Wordpress        = "wordpress"
	Typo3            = "typo3"
	RabbitMQ         = "rabbitmq"
	Spark            = "spark"
	Numpy            = "numpy"
	MeteorJS         = "meteor-js"
	Qt               = "qt"
	Gtk              = "gtk"
	WxWidgets        = "wxWidgets"
	Ansible          = "ansible"
	Chef             = "chef"
)

const (
	MongoDb    = "mongodb"
	MySQL      = "mysql"
	PostgreSQL = "postgresql"
	SQLite     = "sqlite"
	Redis      = "redis"
	Cassandra  = "cassandra"
	Neo4j      = "neo4j"
	OracleDb   = "oracledb"
)

const (
	AmazonAWS        = "amazon-aws"
	AmazonEC2        = "amazon-ec2"
	AmazonCloudTrail = "amazon-cloudtrail"
	AmazonGlacier    = "amazon-glacier"
	AmazonIAM        = "amazon-iam"
	AmazonKMS        = "amazon-kms"
	AmazonRDS        = "amazon-rds"
	AmazonRoute53    = "amazon-route53"
	AmazonS3         = "amazon-s3"
	AmazonSNS        = "amazon-sns"
	AmazonSQS        = "amazon-sqs"
)

const (
	GoogleCloudPlatform        = "google-cloud-platform"
	GoogleCloudBigQuery        = "google-cloud-bigquery"
	GoogleCloudCompute         = "google-cloud-compute"
	GoogleCloudDatastore       = "google-cloud-datastore"
	GoogleCloudDNS             = "google-cloud-dns"
	GoogleCloudErrorReporting  = "google-cloud-errorreporting"
	GoogleCloudLanguage        = "google-cloud-language"
	GoogleCloudLogging         = "google-cloud-logging"
	GoogleCloudMonitoring      = "google-cloud-monitoring"
	GoogleCloudPubSub          = "google-cloud-pubsub"
	GoogleCloudResourceManager = "google-cloud-resourcemanager"
	GoogleCloudSpeech          = "google-cloud-speech"
	GoogleCloudStorage         = "google-cloud-storage"
	GoogleCloudTrace           = "google-cloud-trace"
	GoogleCloudTranslate       = "google-cloud-translate"
	GoogleCloudVision          = "google-cloud-vision"
)
