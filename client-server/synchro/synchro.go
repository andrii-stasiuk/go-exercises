/*Package synchro for synchronization of work*/
package synchro

// Synchronizer channel type for synchronization of work
type Synchronizer chan struct{}

// Stop method of Synchronizer type
func (s Synchronizer) Stop() {
	s <- struct{}{}
}

// Resume method of Synchronizer type
func (s Synchronizer) Resume() {
	<-s
}
