<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: vtgate.proto
//   Date: 2016-01-22 01:34:42

namespace Vitess\Proto\Vtgate\SplitQueryResponse {

  class KeyRangePart extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $keyspace = null;
    
    /**  @var \Vitess\Proto\Topodata\KeyRange[]  */
    public $key_ranges = array();
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'vtgate.SplitQueryResponse.KeyRangePart');

      // OPTIONAL STRING keyspace = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "keyspace";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // REPEATED MESSAGE key_ranges = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "key_ranges";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $f->reference = '\Vitess\Proto\Topodata\KeyRange';
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <keyspace> has a value
     *
     * @return boolean
     */
    public function hasKeyspace(){
      return $this->_has(1);
    }
    
    /**
     * Clear <keyspace> value
     *
     * @return \Vitess\Proto\Vtgate\SplitQueryResponse\KeyRangePart
     */
    public function clearKeyspace(){
      return $this->_clear(1);
    }
    
    /**
     * Get <keyspace> value
     *
     * @return string
     */
    public function getKeyspace(){
      return $this->_get(1);
    }
    
    /**
     * Set <keyspace> value
     *
     * @param string $value
     * @return \Vitess\Proto\Vtgate\SplitQueryResponse\KeyRangePart
     */
    public function setKeyspace( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <key_ranges> has a value
     *
     * @return boolean
     */
    public function hasKeyRanges(){
      return $this->_has(2);
    }
    
    /**
     * Clear <key_ranges> value
     *
     * @return \Vitess\Proto\Vtgate\SplitQueryResponse\KeyRangePart
     */
    public function clearKeyRanges(){
      return $this->_clear(2);
    }
    
    /**
     * Get <key_ranges> value
     *
     * @param int $idx
     * @return \Vitess\Proto\Topodata\KeyRange
     */
    public function getKeyRanges($idx = NULL){
      return $this->_get(2, $idx);
    }
    
    /**
     * Set <key_ranges> value
     *
     * @param \Vitess\Proto\Topodata\KeyRange $value
     * @return \Vitess\Proto\Vtgate\SplitQueryResponse\KeyRangePart
     */
    public function setKeyRanges(\Vitess\Proto\Topodata\KeyRange $value, $idx = NULL){
      return $this->_set(2, $value, $idx);
    }
    
    /**
     * Get all elements of <key_ranges>
     *
     * @return \Vitess\Proto\Topodata\KeyRange[]
     */
    public function getKeyRangesList(){
     return $this->_get(2);
    }
    
    /**
     * Add a new element to <key_ranges>
     *
     * @param \Vitess\Proto\Topodata\KeyRange $value
     * @return \Vitess\Proto\Vtgate\SplitQueryResponse\KeyRangePart
     */
    public function addKeyRanges(\Vitess\Proto\Topodata\KeyRange $value){
     return $this->_add(2, $value);
    }
  }
}

