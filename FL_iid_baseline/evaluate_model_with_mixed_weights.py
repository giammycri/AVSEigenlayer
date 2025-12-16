import warnings
warnings.filterwarnings("ignore")

from Script_FL.generate_random_weights import generate_random_weights

import numpy as np

def evaluate_model_with_mixed_weights(real_weights_list, num_random_sets, model, test_X, test_y):
    """
    Concatena una lista di pesi reali con i pesi randomici, esegue il federated averaging,
    imposta i pesi sul modello e valuta il modello.

    :param real_weights_list: Lista di set di pesi reali da utilizzare.
    :param num_random_sets: Numero di set di pesi randomici da generare. Ovvero il numero di partecipanti malevoli
    :param model: Il modello da valutare.
    :param test_X: I dati di test per il modello.
    :param test_y: Le etichette di test per il modello.
    :return: Il risultato della valutazione del modello.
    """
    # Assumiamo che tutti i pesi nella lista real_weights_list abbiano la stessa forma
    original_weights = real_weights_list[0]

    # Genera pesi randomici basati sulla forma dei pesi originali
    random_weights_sets = generate_random_weights(original_weights, num_random_sets, perturbation_scale=0.1)

    # Concatena i pesi randomici con la lista dei pesi reali
    combined_weights = random_weights_sets + real_weights_list

    # Federated averaging
    sum_list = [np.sum(np.array(x), axis=0) for x in zip(*combined_weights)]
    avg_list = [x / len(combined_weights) for x in sum_list]

    # Imposta i pesi mediati sul modello
    model.set_weights(avg_list)

    # Valuta il modello
    return model.evaluate(test_X, test_y), random_weights_sets
