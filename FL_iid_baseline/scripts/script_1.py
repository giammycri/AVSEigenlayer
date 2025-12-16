
import syft as sy
from keras.datasets import mnist
from syft.service.user.user import ServiceRole
from syft.service.user.user import UserCreate

domain_client = sy.login(port="3331", email="info@openmined.org", password="changethis")
domain_client.users.create(UserCreate(email="alice1@openmined.org", name="Alice1", password="pw", role=ServiceRole.DATA_OWNER))
alice_do_1 = domain_client.login(email="alice1@openmined.org", password="pw")

(train_X_a_1, train_y_a_1), (test_X_a_1, test_y_a_1) = mnist.load_data()

dataset = sy.Dataset(name="MNIST_1")

trainx = sy.Asset(name="Asset di Train")
trainx.set_obj(train_X_a_1)
trainx.no_mock()

trainy = sy.Asset(name="Asset di etichette di Train")
trainy.set_obj(train_y_a_1)
trainy.no_mock()

testx = sy.Asset(name="Asset di Test")
testx.set_obj(test_X_a_1)
testx.no_mock()

testy = sy.Asset(name="Asset di eitchette di Test")
testy.set_obj(test_y_a_1)
testy.no_mock()

dataset.add_asset(trainx)
dataset.add_asset(trainy)
dataset.add_asset(testx)
dataset.add_asset(testy)

alice_do_1.upload_dataset(dataset)  # Aggiunta dell'indice alla chiamata del metodo upload_dataset
