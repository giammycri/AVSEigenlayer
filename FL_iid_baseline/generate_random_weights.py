import warnings
warnings.filterwarnings("ignore")

import numpy as np

def generate_random_weights(original_weights, num_sets, target_norm=0.20):
    """
    Crea dei pesi randomici con la stessa struttura e forma dei pesi originali.
    :param original_weights: Pesi originali.
    :param num_sets: Può essere visto come il numero di partecipanti malevoli.
    :param target_norm: Norma euclidea desiderata per i pesi generati.
    """
    random_weights_sets = []
    last_layer_weights = []

    for _ in range(num_sets):
        random_weights = [np.random.normal(scale=target_norm, size=weights.shape).astype(np.float32) for weights in original_weights]

        random_weights_sets.append(random_weights)

        # Stampa l'ultimo array di pesi randomici generato
        last_layer_weights = random_weights_sets[-1][-1]  # Accede all'ultimo layer dell'ultimo set
        print("Ultimo array dei pesi randomici:", last_layer_weights)

    return random_weights_sets