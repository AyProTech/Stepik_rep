type Strakt struct {
    On   bool
    Ammo  int
    Power int
}
func (a *Strakt) Shoot() bool{
    if a.On==false{
        return false
    }else if a.Ammo>0{
        a.Ammo -= 1
        return true
    }
    return false
}
func (b *Strakt) RideBike() bool{
    if b.On == false{
        return false
    }else if b.Power>0{
        b.Power -=1
        return true
    }
    return false
}
        
func main() {
    testStruct := new(Strakt)