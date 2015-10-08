// Evo is a package to assist the implementation of evolutionary algorithms.
//
// Evo exposes a clean and flexible API oriented around two interfaces: Genome
// and Population. Genomes represent candidate solutions to the user's problem
// and are implemented by the user. Genomes define their own means of evolution,
// allowing for a multiplicity of techniques ranging from genetic algorithms to
// evolution strategies and beyond. Several composeable population types are
// provided by Evo to enable the construction of both common and novel
// architectures.
//
// Genomes define the body of the evolutionary loop as the Evolve method. For
// each genome in a population, the Evolve method is called, passing some subset
// of the population, called the suitors, as arguments. The Evolve method then
// applies operators to the suiters (selection, mutation, etc) and returns a
// genome that will replace the caller within the population for the next
// iteration. Common operators for a variety of representations are provided by
// Evo.
//
// Populations orchestrate the evolution of genomes. Populations provided by Evo
// live under the package evo/pop. The generational population implements a
// traditional generation-based loop with master-slave parallelism. Each genome
// receives the entire population as suitors, and the population is only updated
// after all genomes have returned. The graph population maps each genome to a
// node in a graph. Each genome only receives the neighboring genomes as
// suitors, and each node is evolved in parallel.
//
// Populations themselves implement the Genome interface. The Evolve method on
// populations implements uniform random migration: A random suitor is chosen
// and asserted to be a population of the same type. Then the population and its
// suitor exchange random members. This allows novel architectures like the
// island model to be implemented by nesting populations.
package evo

// TODO: Keep this in sync with the readme
