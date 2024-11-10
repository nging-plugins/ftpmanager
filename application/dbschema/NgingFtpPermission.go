// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/lib/factory/pagination"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingFtpPermission []*NgingFtpPermission

func (s Slice_NgingFtpPermission) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFtpPermission) RangeRaw(fn func(m *NgingFtpPermission) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingFtpPermission) GroupBy(keyField string) map[string][]*NgingFtpPermission {
	r := map[string][]*NgingFtpPermission{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingFtpPermission{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingFtpPermission) KeyBy(keyField string) map[string]*NgingFtpPermission {
	r := map[string]*NgingFtpPermission{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingFtpPermission) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingFtpPermission) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingFtpPermission) FromList(data interface{}) Slice_NgingFtpPermission {
	values, ok := data.([]*NgingFtpPermission)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingFtpPermission{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingFtpPermission(ctx echo.Context) *NgingFtpPermission {
	m := &NgingFtpPermission{}
	m.SetContext(ctx)
	return m
}

// NgingFtpPermission 用户权限
type NgingFtpPermission struct {
	base    factory.Base
	objects []*NgingFtpPermission

	Id         uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	TargetType string `db:"target_type" bson:"target_type" comment:"目标类型" json:"target_type" xml:"target_type"`
	TargetId   uint   `db:"target_id" bson:"target_id" comment:"目标ID" json:"target_id" xml:"target_id"`
	Permission string `db:"permission" bson:"permission" comment:"权限JSON" json:"permission" xml:"permission"`
	Created    uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated    uint   `db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
}

// - base function

func (a *NgingFtpPermission) Trans() factory.Transactioner {
	return a.base.Trans()
}

func (a *NgingFtpPermission) Use(trans factory.Transactioner) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingFtpPermission) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingFtpPermission) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingFtpPermission) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingFtpPermission) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingFtpPermission) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingFtpPermission) ConnID() int {
	return a.base.ConnID()
}

func (a *NgingFtpPermission) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingFtpPermission) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingFtpPermission) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingFtpPermission) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

func (a *NgingFtpPermission) New(structName string, connID ...int) factory.Model {
	return a.base.New(structName, connID...)
}

func (a *NgingFtpPermission) Base_() factory.Baser {
	return &a.base
}

// - current function

func (a *NgingFtpPermission) Objects() []*NgingFtpPermission {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingFtpPermission) XObjects() Slice_NgingFtpPermission {
	return Slice_NgingFtpPermission(a.Objects())
}

func (a *NgingFtpPermission) NewObjects() factory.Ranger {
	return &Slice_NgingFtpPermission{}
}

func (a *NgingFtpPermission) InitObjects() *[]*NgingFtpPermission {
	a.objects = []*NgingFtpPermission{}
	return &a.objects
}

func (a *NgingFtpPermission) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingFtpPermission) Short_() string {
	return "nging_ftp_permission"
}

func (a *NgingFtpPermission) Struct_() string {
	return "NgingFtpPermission"
}

func (a *NgingFtpPermission) Name_() string {
	b := a
	if b == nil {
		b = &NgingFtpPermission{}
	}
	if b.base.Namer() != nil {
		return WithPrefix(b.base.Namer()(b))
	}
	return WithPrefix(factory.TableNamerGet(b.Short_())(b))
}

func (a *NgingFtpPermission) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.SetConnID(source.ConnID())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingFtpPermission) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingFtpPermission) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFtpPermission:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpPermission(*v))
		case []*NgingFtpPermission:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpPermission(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFtpPermission) GroupBy(keyField string, inputRows ...[]*NgingFtpPermission) map[string][]*NgingFtpPermission {
	var rows Slice_NgingFtpPermission
	if len(inputRows) > 0 {
		rows = Slice_NgingFtpPermission(inputRows[0])
	} else {
		rows = Slice_NgingFtpPermission(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingFtpPermission) KeyBy(keyField string, inputRows ...[]*NgingFtpPermission) map[string]*NgingFtpPermission {
	var rows Slice_NgingFtpPermission
	if len(inputRows) > 0 {
		rows = Slice_NgingFtpPermission(inputRows[0])
	} else {
		rows = Slice_NgingFtpPermission(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingFtpPermission) AsKV(keyField string, valueField string, inputRows ...[]*NgingFtpPermission) param.Store {
	var rows Slice_NgingFtpPermission
	if len(inputRows) > 0 {
		rows = Slice_NgingFtpPermission(inputRows[0])
	} else {
		rows = Slice_NgingFtpPermission(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingFtpPermission) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingFtpPermission:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpPermission(*v))
		case []*NgingFtpPermission:
			err = DBI.FireReaded(a, queryParam, Slice_NgingFtpPermission(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingFtpPermission) Insert() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.TargetType) == 0 {
		a.TargetType = "group"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingFtpPermission) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.TargetType) == 0 {
		a.TargetType = "group"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingFtpPermission) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.TargetType) == 0 {
		a.TargetType = "group"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Updatex()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(a).Updatex(); err != nil {
		return
	}
	err = DBI.Fire("updated", a, mw, args...)
	return
}

func (a *NgingFtpPermission) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.TargetType) == 0 {
		a.TargetType = "group"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdateByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).UpdateByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingFtpPermission) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.TargetType) == 0 {
		a.TargetType = "group"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdatexByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).UpdatexByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingFtpPermission) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFtpPermission) UpdatexField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (affected int64, err error) {
	return a.UpdatexFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingFtpPermission) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["target_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["target_type"] = "group"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	editColumns := make([]string, 0, len(kvset))
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingFtpPermission) UpdatexFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (affected int64, err error) {

	if val, ok := kvset["target_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["target_type"] = "group"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Updatex()
	}
	m := *a
	m.FromRow(kvset)
	editColumns := make([]string, 0, len(kvset))
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(kvset).Updatex(); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", &m, editColumns, mw, args...)
	return
}

func (a *NgingFtpPermission) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(keysValues).Update()
	}
	m := *a
	m.FromRow(keysValues.Map())
	if err = DBI.FireUpdate("updating", &m, keysValues.Keys(), mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(keysValues).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, keysValues.Keys(), mw, args...)
}

func (a *NgingFtpPermission) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.TargetType) == 0 {
			a.TargetType = "group"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.TargetType) == 0 {
			a.TargetType = "group"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingFtpPermission) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingFtpPermission) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Deletex()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).Deletex(); err != nil {
		return
	}
	err = DBI.Fire("deleted", a, mw, args...)
	return
}

func (a *NgingFtpPermission) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingFtpPermission) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingFtpPermission) Reset() *NgingFtpPermission {
	a.Id = 0
	a.TargetType = ``
	a.TargetId = 0
	a.Permission = ``
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *NgingFtpPermission) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["TargetType"] = a.TargetType
		r["TargetId"] = a.TargetId
		r["Permission"] = a.Permission
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "TargetType":
			r["TargetType"] = a.TargetType
		case "TargetId":
			r["TargetId"] = a.TargetId
		case "Permission":
			r["Permission"] = a.Permission
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingFtpPermission) FromRow(row map[string]interface{}) {
	for key, value := range row {
		if _, ok := value.(db.RawValue); ok {
			continue
		}
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "target_type":
			a.TargetType = param.AsString(value)
		case "target_id":
			a.TargetId = param.AsUint(value)
		case "permission":
			a.Permission = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *NgingFtpPermission) GetField(field string) interface{} {
	switch field {
	case "Id":
		return a.Id
	case "TargetType":
		return a.TargetType
	case "TargetId":
		return a.TargetId
	case "Permission":
		return a.Permission
	case "Created":
		return a.Created
	case "Updated":
		return a.Updated
	default:
		return nil
	}
}

func (a *NgingFtpPermission) GetAllFieldNames() []string {
	return []string{
		"Id",
		"TargetType",
		"TargetId",
		"Permission",
		"Created",
		"Updated",
	}
}

func (a *NgingFtpPermission) HasField(field string) bool {
	switch field {
	case "Id":
		return true
	case "TargetType":
		return true
	case "TargetId":
		return true
	case "Permission":
		return true
	case "Created":
		return true
	case "Updated":
		return true
	default:
		return false
	}
}

func (a *NgingFtpPermission) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "TargetType":
			a.TargetType = param.AsString(vv)
		case "TargetId":
			a.TargetId = param.AsUint(vv)
		case "Permission":
			a.Permission = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *NgingFtpPermission) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["target_type"] = a.TargetType
		r["target_id"] = a.TargetId
		r["permission"] = a.Permission
		r["created"] = a.Created
		r["updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "target_type":
			r["target_type"] = a.TargetType
		case "target_id":
			r["target_id"] = a.TargetId
		case "permission":
			r["permission"] = a.Permission
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingFtpPermission) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	return pagination.ListPage(a, cond, sorts...)
}

func (a *NgingFtpPermission) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	return pagination.ListPageAs(a, recv, cond, sorts...)
}

func (a *NgingFtpPermission) ListPageByOffset(cond *db.Compounds, sorts ...interface{}) error {
	return pagination.ListPageByOffset(a, cond, sorts...)
}

func (a *NgingFtpPermission) ListPageByOffsetAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	return pagination.ListPageByOffsetAs(a, recv, cond, sorts...)
}

func (a *NgingFtpPermission) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingFtpPermission) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
