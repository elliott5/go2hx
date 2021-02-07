package stdgo;
import haxe.Rest;
import haxe.ds.Vector;
import stdgo.GoArray;
abstract Slice<T>(SliceData<T>) {

    public var length(get,never):Int;
    private function get_length():Int {
        return this.length;
    }
    private function set_length(length:Int):Int {
        return 0;
    }
    public function new(length:Int,cap:Int=0,args:Rest<T>) {
        this = new SliceData(length,cap);
        if (args.length == 0)
            return;
        for (i in 0...args.length) {
            set(i,args[i]);
        }
    }
    @:op([]) inline function get(index:Int):T {
        return this.get(index);
    }
    @:op([]) inline function set(index:Int,value:T):T {
        return this.set(index,value);
    }
    inline public function slice(low:Int,high:Int=-1):Slice<T> {
        var pos = low;
        if (high == -1)
            high = length;
        var length = high - low;
        var slice = new Slice<T>(0);
        slice.setUnderlying(this.underlying(),pos,length);
        return slice; 
    }
    public function copy():Slice<T> {
        var slice = new Slice<T>(this.length,this.cap());
        for (i in 0...slice.length) {
            slice[i] = 
            get(i);
        }
        return slice;
    }
    public function append(args:Array<T>):Slice<T> {
        var slice = copy();
        var pos = slice.length;
        slice.grow(args.length);
        for (i in 0...args.length) {
            slice[pos + i] = args[i];
        }
        return slice;
    }
    inline public function grow(size:Int) {
        this.grow(size);
    }
    inline public function iterator():Iterator<T> {
        return this.iterator();
    }
    inline public function cap():Int {
        return this.cap();
    }
    inline public function toArray() {
        return this.toArray();
    }
    inline public function toVector() {
        return this.toVector();
    }
    inline public function setUnderlying(array:GoArray<T>,pos:Int,length:Int) {
        this.setUnderlying(array,pos,length);
    }
}
class SliceIterator<T> {
    var pos:Int = 0;
    var slice:SliceData<T>;
    public inline function new(slice) {
        this.slice = slice;
    }
    public inline function hasNext() {
        return pos < slice.length;
    }
    public inline function next() {
        return slice.get(pos++);
    }
}
class SliceData<T> {
    var array:GoArray<T>;
    public var pos:Int = 0;
    public var length:Int = 0;
    public static var forceLength:Int = 0;
    public function new(length:Int=0,cap:Int=0) {
        this.length = length;
        if (cap == 0)
            cap = length;
        array = new GoArray<T>(cap);
    }
    public function get(index:Int):T {
        return array.get(index + pos);
    }
    public function set(index:Int,value:T):T {
        return array.set(index + pos,value);
    }
    inline public function iterator():Iterator<T> {
        return new SliceIterator(this);
    }
    inline public function cap():Int {
        return 0;
    }
    inline public function toArray():Array<T> { //unrolling derefrences
        return [for (i in 0...length) get(i)];
    }
    inline public function toVector():Vector<T> { //derefrence
        var vector = array.toVector();
        var dest = new Vector<T>(length);
        Vector.blit(vector,pos,dest,0,length);
        return dest;
    }
    inline public function toString():String {
        return toArray().toString();
    }
    inline public function underlying():GoArray<T> {
        return array;
    }
    inline public function grow(size:Int) {
        if (array.length >= size + length)
            return;
        length += size;
        var dest = new Vector<T>(length);
        Vector.blit(array.toVector(),0,dest,0,array.length);
        array.replace(dest);
    }
    inline public function setUnderlying(value:GoArray<T>,pos:Int,length:Int) {
        array = value;
        this.pos = pos;
        this.length = length;
    }
}