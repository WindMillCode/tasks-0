import { ENV } from "@env/environment.dev";
import { LinkedList, arraysEqual, clearArray, convertMilitaryToStandard, deepCopy, documentQuerySelector, documentQuerySelectorAll, getQueryPropByName, isNetworkConnectionGood, isPlainObject, numberParse, retriveValueFromPXUnit, setColorBasedOnHEXBackgroundColor, toggleDarkMode, transformObjectKeys, transformPropertiesOnObject, triggerEvent, updateWebStorage } from "./common-utils";

describe("CommonUtils",()=>{
  describe('isNetworkConnectionGood', () => {


    it('should return true when navigator is online', () => {
      // arrange
      spyOnProperty(navigator, 'onLine').and.returnValue(true);

      // act
      const result = isNetworkConnectionGood();

      // assert
      expect(result).toBeTrue();
    });

    it('should return false when navigator is offline', () => {
      // arrange
      spyOnProperty(navigator, 'onLine').and.returnValue(false);

      // act
      const result = isNetworkConnectionGood();

      // assert
      expect(result).toBeFalse();
    });
  });

  describe('triggerEvent', () => {
    beforeEach(() => {
      // any setup you might need
    });

    it('should trigger a custom event successfully', () => {
      // arrange
      const element = document.createElement('div');
      const eventName = 'customEvent';
      let eventTriggered = false;

      // act
      element.addEventListener(eventName, () => {
        eventTriggered = true;
      });
      triggerEvent(element, eventName);

      // assert
      expect(eventTriggered).toBe(true);
    });

    it('should trigger a keydown event successfully', () => {
      // arrange
      const element = document.createElement('div');
      const eventName = 'keydown';
      const keyCode = 13;
      let eventTriggered = false;

      // act
      element.addEventListener(eventName, (event) => {
        eventTriggered = true;
        // assertion for keyCode
        expect(event.keyCode).toBe(keyCode);
      });
      triggerEvent(element, eventName, keyCode);

      // assert
      expect(eventTriggered).toBe(true);
    });

    it('should trigger a keyup event successfully', () => {
      // arrange
      const element = document.createElement('div');
      const eventName = 'keyup';
      const keyCode = 13;
      let eventTriggered = false;

      // act
      element.addEventListener(eventName, (event) => {
        eventTriggered = true;
        // assertion for keyCode
        expect(event.keyCode).toBe(keyCode);
      });
      triggerEvent(element, eventName, keyCode);

      // assert
      expect(eventTriggered).toBe(true);
    });

    it('should fallback to legacy event creation when an exception occurs', () => {
      // Arrange
      const eventName = 'click';
      let receivedEvent: Event | null = null;
      spyOn(document, 'createEvent').and.callThrough();
      let element = document.createElement('div');
      element.addEventListener(eventName, (event) => {
        receivedEvent = event;
      });
      let counter = 0
      element.dispatchEvent = (event) => {
        if(counter===0){
          counter++
          throw new Error('error')
        }
        return true
      };

      // Act
      triggerEvent(element, eventName);

      // Assert
      expect(document.createEvent).toHaveBeenCalled();
      expect(receivedEvent).toBeNull();
    });


  });

  describe('numberParse', () => {
    beforeEach(() => {
      // any setup you might need
    });

    it('should parse a single string dimension correctly', () => {
      // arrange
      const dimensionString = '20px';

      // act
      const result = numberParse(dimensionString);

      // assert
      expect(result).toBe(20);
    });

    it('should parse an array of string dimensions correctly', () => {
      // arrange
      const dimensionArray = ['30px', '40px', '50px'];

      // act
      const result = numberParse(dimensionArray);

      // assert
      // @ts-ignore
      expect(result).toEqual([30, 40, 50]);
    });








  });

  describe('clearArray', () => {
    beforeEach(() => {
      // any setup you might need
    });

    it('should clear the array', () => {
      // arrange
      const arrayToClear = [1, 2, 3, 4, 5];

      // act
      clearArray(arrayToClear);

      // assert
      expect(arrayToClear.length).toBe(0);
    });

    it('should work correctly with an empty array', () => {
      // arrange
      const emptyArray: any[] = [];

      // act
      clearArray(emptyArray);

      // assert
      expect(emptyArray.length).toBe(0);
    });


  });

  describe('getQueryPropByName', () => {
    beforeEach(() => {
      // arrange
      // any setup you might need
    });

    it('should return the value of the specified query propeter', () => {
      // arrange
      const propName = 'example';
      const propValue = '123';
      const testUrl = `https://example.com/?${propName}=${propValue}`;

      // act
      const result = getQueryPropByName(propName, testUrl);

      // assert
      expect(result).toBe(propValue);
    });

    it('should return null for a non-existing query propeter', () => {
      // arrange
      const nonExistingProp = 'nonexistent';
      const testUrl = `https://example.com/`;

      // act
      const result = getQueryPropByName(nonExistingProp, testUrl);

      // assert
      expect(result).toEqual(null);
    });

  });

  describe('documentQuerySelector', () => {
    it(` when called |
        with a valid selector |
        returns the HTMLElement matching the selector`, () => {
      // arrange
      const validSelector = 'body';

      // act
      const result = documentQuerySelector(validSelector);

      // assert
      expect(result).toBeInstanceOf(HTMLElement);

    });

    it(` when called |
        with an invalid selector |
        returns null`, () => {
      // arrange
      const invalidSelector = '#nonExistentElement';

      // act
      const result = documentQuerySelector(invalidSelector);

      // assert
      expect(result).toBeNull();

    });
  });

  describe('documentQuerySelectorAll', () => {
    it(` when called |
        with a valid selector |
        returns an array of HTMLElements matching the selector`, () => {
      // arrange
      const validSelector = '.myClass';

      // act
      const result = documentQuerySelectorAll(validSelector);

      // assert
      expect(result).toBeInstanceOf(Array);
      expect(result.every((item) => item instanceof HTMLElement)).toBeTrue();

    });

    it(` when called |
        with an invalid selector |
        returns an empty array`, () => {
      // arrange
      const invalidSelector = '.nonExistentClass';

      // act
      const result = documentQuerySelectorAll(invalidSelector);

      // assert
      expect(result).toEqual([]);

    });
  });

  describe('deepCopy', () => {
    it(` when called |
        with a valid object |
        returns a deep copy of the object`, () => {
      // arrange
      const originalObject = { key: 'value', nested: { innerKey: 'innerValue' } };

      // act
      const result = deepCopy(originalObject);

      // assert
      expect(result).toEqual(originalObject);
      expect(result).not.toBe(originalObject);

    });
  });

  describe('isPlainObject', () => {

    it('should return true for a plain object', () => {
      // arrange
      const plainObject = { key: 'value' };

      // act
      const result = isPlainObject(plainObject);

      // assert
      expect(result).toBeTrue();
    });

    it('should return false for null', () => {
      // arrange
      const nullValue = null;

      // act
      const result = isPlainObject(nullValue);

      // assert
      expect(result).toBeFalse();
    });

    it('should return false for undefined', () => {
      // arrange
      const undefinedValue = undefined;

      // act
      const result = isPlainObject(undefinedValue);

      // assert
      expect(result).toBeFalse();
    });

    it('should return false for non-object values', () => {
      // arrange
      const nonObjectValues = [42, 'string', true, Symbol('symbol')];

      // act
      const results = nonObjectValues.map(value => isPlainObject(value));

      // assert
      results.forEach(result => expect(result).toBeFalse());
    });

    it('should return false for arrays', () => {
      // arrange
      const array = [1, 2, 3];

      // act
      const result = isPlainObject(array);

      // assert
      expect(result).toBeFalse();
    });

    it('should return false for functions', () => {
      // arrange
      const func = function () {};

      // act
      const result = isPlainObject(func);

      // assert
      expect(result).toBeFalse();
    });

    it('should return false for instances of classes', () => {
      // arrange
      class TestClass {}

      const instance = new TestClass();

      // act
      const result = isPlainObject(instance);

      // assert
      expect(result).toBeFalse();
    });

    it('should return true for a plain object', () => {
      // Arrange
      const plainObject = { key: 'value' };

      // Act
      const result = isPlainObject(plainObject);

      // Assert
      expect(result).toBeTrue();
    });

    it('should return true for an object with a null prototype', () => {
      // Arrange
      const objWithNullProto = Object.create(null);

      // Act
      const result = isPlainObject(objWithNullProto);

      // Assert
      expect(result).toBeTrue();
    });

    it('should return false for null', () => {
      // Arrange & Act
      const result = isPlainObject(null);

      // Assert
      expect(result).toBeFalse();
    });

    it('should return false for undefined', () => {
      // Arrange & Act
      const result = isPlainObject(undefined);

      // Assert
      expect(result).toBeFalse();
    });

    it('should return false for non-object values', () => {
      // Arrange
      const values = [42, 'string', true, Symbol('symbol'), function() {}, []];

      // Act & Assert
      values.forEach(value => {
        expect(isPlainObject(value)).toBeFalse();
      });
    });

    it('should return false for an instance of a class', () => {
      // Arrange
      class MyClass {}
      const instance = new MyClass();

      // Act
      const result = isPlainObject(instance);

      // Assert
      expect(result).toBeFalse();
    });



  });

  describe('transformObjectKeys', () => {

    it('should transform keys of a plain object using the provided predicate', () => {
      // arrange
      const inputObject = { myKey: 'myValue', nested: { innerKey: 'innerValue' } };
      const predicate = key => key.toUpperCase();

      // act
      const result = transformObjectKeys(inputObject, predicate);

      // assert
      expect(result).toEqual({ MYKEY: 'myValue', NESTED: { INNERKEY: 'innerValue' } });
    });

    it('should handle an empty object', () => {
      // arrange
      const inputObject = {};
      const predicate = key => key.toUpperCase();

      // act
      const result = transformObjectKeys(inputObject, predicate);

      // assert
      expect(result).toEqual({});
    });

    it('should handle an object with different value types', () => {
      // arrange
      const inputObject = { stringKey: 'stringValue', numberKey: 42, boolKey: true, arrayKey: [1, 2, 3] };
      const predicate = key => key.toUpperCase();

      // act
      const result = transformObjectKeys(inputObject, predicate);

      // assert
      expect(result).toEqual({
        STRINGKEY: 'stringValue',
        NUMBERKEY: 42,
        BOOLKEY: true,
        ARRAYKEY: [1, 2, 3]
      });
    });

    it('should handle an array of objects', () => {
      // arrange
      const inputArray = [
        { key1: 'value1', key2: 'value2' },
        { key3: 'value3', key4: 'value4' }
      ];
      const predicate = key => key.toUpperCase();

      // act
      const result = transformObjectKeys(inputArray, predicate);

      // assert
      expect(result).toEqual([
        { KEY1: 'value1', KEY2: 'value2' },
        { KEY3: 'value3', KEY4: 'value4' }
      ]);
    });

    it('should handle nested arrays and objects', () => {
      // arrange
      const inputObject = {
        key1: 'value1',
        key2: {
          key3: 'value3',
          key4: [{ key5: 'value5' }, { key6: 'value6' }]
        }
      };
      const predicate = key => key.toUpperCase();

      // act
      const result = transformObjectKeys(inputObject, predicate);

      // assert
      expect(result).toEqual({
        KEY1: 'value1',
        KEY2: {
          KEY3: 'value3',
          KEY4: [{ KEY5: 'value5' }, { KEY6: 'value6' }]
        }
      });
    });

  });

  describe('retriveValueFromPXUnit', () => {

    it('should retrieve the numeric value from a string with "px" unit', () => {
      // arrange
      const inputValue = '10px';

      // act
      const result = retriveValueFromPXUnit(inputValue);

      // assert
      expect(result).toBe('10');
    });

    it('should handle values with multiple digits', () => {
      // arrange
      const inputValue = '12345px';

      // act
      const result = retriveValueFromPXUnit(inputValue);

      // assert
      expect(result).toBe('12345');
    });

    it('should handle values with leading zeros', () => {
      // arrange
      const inputValue = '005px';

      // act
      const result = retriveValueFromPXUnit(inputValue);

      // assert
      expect(result).toBe('005');
    });

    it('should handle values without leading zeros', () => {
      // arrange
      const inputValue = '50px';

      // act
      const result = retriveValueFromPXUnit(inputValue);

      // assert
      expect(result).toBe('50');
    });

    it('should handle values with decimal points', () => {
      // arrange
      const inputValue = '12.34px';

      // act
      const result = retriveValueFromPXUnit(inputValue);

      // assert
      expect(result).toBe('12');
    });

    it('should handle values without "px" unit', () => {
      // arrange
      const inputValue = '100';

      // act
      const result = retriveValueFromPXUnit(inputValue);

      // assert
      expect(result).toBe('100');
    });

    it('should handle values with no numeric part', () => {
      // arrange
      const inputValue = 'px';

      // act
      const result = retriveValueFromPXUnit(inputValue);

      // assert
      expect(result).toBeUndefined();
    });

    it('should handle values with negative sign', () => {
      // arrange
      const inputValue = '-15px';

      // act
      const result = retriveValueFromPXUnit(inputValue);

      // assert
      expect(result).toBe('-15');
    });

    it('should handle values with spaces', () => {
      // arrange
      const inputValue = ' 25px ';

      // act
      const result = retriveValueFromPXUnit(inputValue);

      // assert
      expect(result).toBe('25');
    });

  });

  describe('setColorBasedOnHEXBackgroundColor', () => {

    it('should return "white" for a dark background color', () => {
      // arrange
      const darkColor = '#222222';

      // act
      const result = setColorBasedOnHEXBackgroundColor(darkColor);

      // assert
      expect(result).toEqual('white');
    });

    it('should return "black" for a light background color', () => {
      // arrange
      const lightColor = '#eeeeee';

      // act
      const result = setColorBasedOnHEXBackgroundColor(lightColor);

      // assert
      expect(result).toEqual('black');
    });

    it('should return "black" for a transparent background color', () => {
      // arrange
      const transparentColor = 'transparent';

      // act
      const result = setColorBasedOnHEXBackgroundColor(transparentColor);

      // assert
      expect(result).toEqual('black');
    });

    it('should return "black" for an invalid color format', () => {
      // arrange
      const invalidColor = 'invalidColor';

      // act
      const result = setColorBasedOnHEXBackgroundColor(invalidColor);

      // assert
      expect(result).toEqual('black');
    });



  });

  describe('convertMilitaryToStandard', () => {

    it('should convert military time to standard time (PM)', () => {
      // arrange
      const militaryTime = '18:30';

      // act
      const result = convertMilitaryToStandard(militaryTime);

      // assert
      expect(result).toEqual('6:30 PM');
    });

    it('should convert military time to standard time (AM)', () => {
      // arrange
      const militaryTime = '08:45';

      // act
      const result = convertMilitaryToStandard(militaryTime);

      // assert
      expect(result).toEqual('08:45 AM');
    });

    it('should convert 12:00 to "12:00 PM"', () => {
      // arrange
      const militaryTime = '12:00';

      // act
      const result = convertMilitaryToStandard(militaryTime);

      // assert
      expect(result).toEqual('12:00 PM');
    });

    it('should convert 00:00 to "12:00 AM"', () => {
      // arrange
      const militaryTime = '00:00';

      // act
      const result = convertMilitaryToStandard(militaryTime);

      // assert
      expect(result).toEqual('12:00 AM');
    });

    it('should handle invalid input', () => {
      // arrange
      const invalidTime = 'invalidTime';

      // act
      expect(()=>convertMilitaryToStandard(invalidTime)).toThrow();

      // assert
    });



  });

  describe('transformPropertiesOnObject', () => {

    it('should transform properties on a simple object', () => {
      // arrange
      const data = { firstName: 'John', lastName: 'Doe' };
      const predicate = key => key.toUpperCase();

      // act
      const result = transformPropertiesOnObject(data, predicate);

      // assert
      expect(result).toEqual({ FIRSTNAME: 'John', LASTNAME: 'Doe' });
    });

    it('should transform properties on an object with nested objects', () => {
      // arrange
      const data = {
        person: {
          firstName: 'John',
          lastName: 'Doe',
          address: {
            city: 'New York',
            country: 'USA'
          }
        }
      };
      const predicate = key => key.toUpperCase();

      // act
      const result = transformPropertiesOnObject(data, predicate);

      // assert
      expect(result).toEqual({
        PERSON: {
          FIRSTNAME: 'John',
          LASTNAME: 'Doe',
          ADDRESS: {
            CITY: 'New York',
            COUNTRY: 'USA'
          }
        }
      });
    });

    it('should transform properties on an array of objects', () => {
      // arrange
      const data = [
        { name: 'Apple', color: 'Red' },
        { name: 'Banana', color: 'Yellow' }
      ];
      const predicate = key => key.toUpperCase();

      // act
      const result = transformPropertiesOnObject(data, predicate);

      // assert
      expect(result).toEqual([
        { NAME: 'Apple', COLOR: 'Red' },
        { NAME: 'Banana', COLOR: 'Yellow' }
      ]);
    });

    it('should handle non-object values', () => {
      // arrange
      const data = 'test';
      const predicate = key => key.toUpperCase();

      // act
      const result = transformPropertiesOnObject(data, predicate);

      // assert
      expect(result).toEqual('test');
    });



  });

 describe('updateWebStorage', () => {

    let webStorage: Storage;
    let storageTitle: string;
    let predicate: jasmine.Spy;

    beforeEach(() => {
      // arrange
      webStorage = window.sessionStorage;
      storageTitle = 'testStorage';
      spyOn(webStorage,"getItem").and.callThrough()
      predicate = jasmine.createSpy('predicate').and.returnValue({ updated: true });
    });

    it('should update web storage using the provided predicate', () => {
      // arrange
      (webStorage.getItem as jasmine.Spy).and.returnValue(JSON.stringify({ existingData: true }));

      // act
      const result = updateWebStorage(webStorage, storageTitle, predicate);

      // assert
      expect(webStorage.getItem).toHaveBeenCalledWith(storageTitle);
      expect(predicate).toHaveBeenCalledWith({ existingData: true });
      expect(result).toEqual({ updated: true });
    });

    it('should handle empty web storage', () => {
      // act
      const result = updateWebStorage(webStorage, storageTitle, predicate);

      // assert
      expect(webStorage.getItem).toHaveBeenCalledWith(storageTitle);
      expect(predicate).toHaveBeenCalledWith({});
      expect(result).toEqual({ updated: true });
    });

    it('should handle non-existent web storage item', () => {
      // arrange
      (webStorage.getItem as jasmine.Spy).and.returnValue(null);

      // act
      const result = updateWebStorage(webStorage, storageTitle, predicate);

      // assert
      expect(webStorage.getItem).toHaveBeenCalledWith(storageTitle);
      expect(predicate).toHaveBeenCalledWith({});
      expect(result).toEqual({ updated: true });
    });



  });


  describe('toggleDarkMode', () => {

    let root = document.querySelector(':root') as HTMLElement;
    let localStorageSpy: jasmine.SpyObj<any>;
    let init: boolean;
    let colorMode: "light" | "dark" | undefined;

    beforeEach(() => {
      // arrange
      let obj = JSON.stringify({
        [ENV.classPrefix.app]:"{}"
      })

      spyOn(localStorage, 'getItem').and.callFake((key) => {return JSON.parse(obj)[key]})
      spyOn(localStorage, 'setItem').and.callFake((key, value) => {

        let newObj = JSON.parse(obj)
        newObj[key]=value
        obj = JSON.stringify(newObj)
      });
      init = false;
      colorMode = undefined;
    });

    it('should toggle dark mode when init is false and no color mode provided', () => {
      // arrange

      // act
      toggleDarkMode(init, colorMode);

      // assert
      expect(localStorage.getItem).toHaveBeenCalledWith(ENV.classPrefix.app);
      expect(localStorage.setItem).toHaveBeenCalledWith(
        ENV.classPrefix.app,
        JSON.stringify({ darkMode: true })
      );
      expect(root.classList.contains('WMLDarkMode')).toBeTrue();
      expect(root.classList.contains('WMLLightMode')).toBeFalse();
    });

    it('should toggle dark mode when init is false and color mode is light', () => {
      // arrange

      // act
      toggleDarkMode(init, 'light');

      // assert
      expect(localStorage.getItem).toHaveBeenCalledWith(ENV.classPrefix.app);
      expect(localStorage.setItem).toHaveBeenCalledWith(
        ENV.classPrefix.app,
        JSON.stringify({ darkMode: false })
      );
      expect(root.classList.contains('WMLLightMode')).toBeTrue();
      expect(root.classList.contains('WMLDarkMode')).toBeFalse();
    });

    it('should toggle dark mode when init is false and color mode is dark', () => {
      // arrange

      // act
      toggleDarkMode(init, 'dark');

      // assert
      expect(localStorage.getItem).toHaveBeenCalledWith(ENV.classPrefix.app);
      expect(localStorage.setItem).toHaveBeenCalledWith(
        ENV.classPrefix.app,
        JSON.stringify({ darkMode: true })
      );
      expect(root.classList.contains('WMLDarkMode')).toBeTrue();
      expect(root.classList.contains('WMLLightMode')).toBeFalse();
    });

    it('should initialize dark mode when init is true and dark mode is undefined', () => {
      // arrange

      // act
      toggleDarkMode(true);

      // assert
      expect(localStorage.getItem).toHaveBeenCalledWith(ENV.classPrefix.app);
      expect(localStorage.setItem).toHaveBeenCalledWith(
        ENV.classPrefix.app,
        JSON.stringify({})
      );
      expect(root.classList.contains('WMLDarkMode')).toEqual(false);
      expect(root.classList.contains('WMLLightMode')).toEqual(true);
    });



  });

  describe('arraysEqual', () => {

    it('should return true for equal arrays', () => {
      // arrange
      const a = [1, 'hello', { key: 'value' }];
      const b = ['hello', { key: 'value' }, 1];

      // act
      const result = arraysEqual(a, b);

      // assert
      expect(result).toEqual(true);
    });

    it('should return false for arrays with different lengths', () => {
      // arrange
      const a = [1, 2, 3];
      const b = [1, 2];

      // act
      const result = arraysEqual(a, b);

      // assert
      expect(result).toEqual(false);
    });

    it('should return false for arrays with different elements', () => {
      // arrange
      const a = [1, 'hello', { key: 'value' }];
      const b = ['hello', { key: 'differentValue' }, 1];

      // act
      const result = arraysEqual(a, b);

      // assert
      expect(result).toEqual(false);
    });

    it('should handle arrays with repeated elements', () => {
      // arrange
      const a = [1, 2, 3, 3, 'hello'];
      const b = [1, 'hello', 3, 2, 3];

      // act
      const result = arraysEqual(a, b);

      // assert
      expect(result).toEqual(true);
    });



  });


describe('LinkedList', () => {
  let linkedList: LinkedList<number>;

  beforeEach(() => {
    linkedList = new LinkedList<number>(1);
  });

  describe('init', () => {
    it('should create', () => {
      expect(linkedList).toBeTruthy();
    });

    it('should have head initialized properly', () => {
      const head = linkedList.getHead();
      expect(head.val).toEqual(1);
      expect(head.next).toBeNull();
    });
  });

  describe('addNode', () => {
    it('should add a node to the list', () => {
      // arrange
      linkedList.addNode(2);

      // act
      const head = linkedList.getHead();

      // assert
      expect(head.val).toEqual(1);
      expect(head.next.val).toEqual(2);
      expect(head.next.next).toBeNull();
    });
  });

  describe('addArrayItemsToList', () => {
    it('should add array items to the list', () => {
      // arrange
      const arrayItems = [2, 3, 4];

      // act
      linkedList.addArrayItemsToList(arrayItems);

      // assert
      const head = linkedList.getHead();
      expect(head.val).toEqual(1);
      expect(head.next.val).toEqual(2);
      expect(head.next.next.val).toEqual(3);
      expect(head.next.next.next.val).toEqual(4);
      expect(head.next.next.next.next).toBeNull();
    });
  });

  describe('moveToNextItemInList', () => {
    it('should move to the next item in the list', () => {
      // arrange
      linkedList.addNode(2);
      linkedList.addNode(3);

      // act
      linkedList.moveToNextItemInList();

      // assert
      const head = linkedList.getHead();
      expect(head.val).toEqual(1);
      expect(head.next.val).toEqual(2);
      expect(head.next.next).toEqual({ val: 3, next: null });
    });
  });

  describe('closeList', () => {
    it('should close the list by connecting the last node to the head', () => {
      // arrange
      linkedList.addNode(2);
      linkedList.addNode(3);
      // act
      linkedList.closeList();

      // assert
      const head = linkedList.getHead();
      expect(head.val).toEqual(1);
      expect(head.next.val).toEqual(2);
      expect(head.next.next.val).toEqual(3);
      expect(head.next.next.next).toEqual(head);
    });
  });
});

})
