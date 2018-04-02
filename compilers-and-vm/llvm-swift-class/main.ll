; ModuleID = '-'
source_filename = "-"
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

%Ts6UInt32V = type <{ i32 }>
%swift.refcounted = type { %swift.type*, i32, i32 }
%swift.type = type { i64 }
%T4main13SimpleStorageC = type <{ %swift.refcounted }>
%swift.opaque = type opaque
%swift.type_metadata_record = type { i32, i32 }

@_T0s19_emptyStringStorages6UInt32Vv = external global %Ts6UInt32V, align 4
@0 = private unnamed_addr constant [14 x i8] c"Hello, World!\00"
@_swift_allocObject = external global %swift.refcounted* (%swift.type*, i64, i64)*
@_T0BoWV = external global i8*, align 8
@1 = private constant [22 x i8] c"4main13SimpleStorageC\00"
@2 = private constant [1 x i8] zeroinitializer
@_T04main13SimpleStorageCMn = hidden constant <{ i32, i32, i32, i32, i32, i32, i32, i32, i32, i32, i32 }> <{ i32 trunc (i64 sub (i64 ptrtoint ([22 x i8]* @1 to i64), i64 ptrtoint (<{ i32, i32, i32, i32, i32, i32, i32, i32, i32, i32, i32 }>* @_T04main13SimpleStorageCMn to i64)) to i32), i32 0, i32 12, i32 trunc (i64 sub (i64 ptrtoint ([1 x i8]* @2 to i64), i64 ptrtoint (i32* getelementptr inbounds (<{ i32, i32, i32, i32, i32, i32, i32, i32, i32, i32, i32 }>, <{ i32, i32, i32, i32, i32, i32, i32, i32, i32, i32, i32 }>* @_T04main13SimpleStorageCMn, i32 0, i32 3) to i64)) to i32), i32 trunc (i64 sub (i64 ptrtoint (%swift.type** (%swift.type*)* @get_field_types_SimpleStorage to i64), i64 ptrtoint (i32* getelementptr inbounds (<{ i32, i32, i32, i32, i32, i32, i32, i32, i32, i32, i32 }>, <{ i32, i32, i32, i32, i32, i32, i32, i32, i32, i32, i32 }>* @_T04main13SimpleStorageCMn, i32 0, i32 4) to i64)) to i32), i32 0, i32 trunc (i64 sub (i64 ptrtoint (%swift.type* ()* @_T04main13SimpleStorageCMa to i64), i64 ptrtoint (i32* getelementptr inbounds (<{ i32, i32, i32, i32, i32, i32, i32, i32, i32, i32, i32 }>, <{ i32, i32, i32, i32, i32, i32, i32, i32, i32, i32, i32 }>* @_T04main13SimpleStorageCMn, i32 0, i32 6) to i64)) to i32), i32 10, i32 0, i32 0, i32 0 }>, section ".rodata", align 8
@_T04main13SimpleStorageCML = internal global %swift.type* null, align 8
@_T04main13SimpleStorageCMf = internal global <{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }> <{ void (%T4main13SimpleStorageC*)* @_T04main13SimpleStorageCfD, i8** @_T0BoWV, i64 0, %swift.type* null, %swift.opaque* null, %swift.opaque* null, i64 1, i32 3, i32 0, i32 16, i16 7, i16 0, i32 112, i32 16, i64 sub (i64 ptrtoint (<{ i32, i32, i32, i32, i32, i32, i32, i32, i32, i32, i32 }>* @_T04main13SimpleStorageCMn to i64), i64 ptrtoint (i64* getelementptr inbounds (<{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>, <{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>* @_T04main13SimpleStorageCMf, i32 0, i32 14) to i64)), i8* null, { i64, i64, i64 } (%T4main13SimpleStorageC*)* @_T04main13SimpleStorageC10printHelloSSyF, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* @_T04main13SimpleStorageCACycfc }>, align 8
@3 = private constant [22 x i8] c"4main13SimpleStorageC\00", section ".swift3_typeref"
@_T04main13SimpleStorageCMF = internal constant { i32, i32, i16, i16, i32 } { i32 trunc (i64 sub (i64 ptrtoint ([22 x i8]* @3 to i64), i64 ptrtoint ({ i32, i32, i16, i16, i32 }* @_T04main13SimpleStorageCMF to i64)) to i32), i32 0, i16 1, i16 12, i32 0 }, section ".swift3_fieldmd", align 4
@field_type_vector_SimpleStorage = private global %swift.type** null
@_swift_slowAlloc = external global i8* (i64, i64)*
@_swift_slowDealloc = external global void (i8*, i64, i64)*
@"\01l_type_metadata_table" = private constant [1 x %swift.type_metadata_record] [%swift.type_metadata_record { i32 trunc (i64 sub (i64 ptrtoint (i64* getelementptr inbounds (<{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>, <{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>* @_T04main13SimpleStorageCMf, i32 0, i32 2) to i64), i64 ptrtoint ([1 x %swift.type_metadata_record]* @"\01l_type_metadata_table" to i64)) to i32), i32 15 }], section ".swift2_type_metadata", align 8
@__swift_reflection_version = linkonce_odr hidden constant i16 3
@_swift1_autolink_entries = private constant [37 x i8] c"-lswiftCore\00-lswiftSwiftOnoneSupport\00", section ".swift1_autolink_entries", align 8
@llvm.used = appending global [4 x i8*] [i8* bitcast ({ i32, i32, i16, i16, i32 }* @_T04main13SimpleStorageCMF to i8*), i8* bitcast ([1 x %swift.type_metadata_record]* @"\01l_type_metadata_table" to i8*), i8* bitcast (i16* @__swift_reflection_version to i8*), i8* getelementptr inbounds ([37 x i8], [37 x i8]* @_swift1_autolink_entries, i32 0, i32 0)], section "llvm.metadata", align 8

@_T04main13SimpleStorageCN = hidden alias %swift.type, bitcast (i64* getelementptr inbounds (<{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>, <{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>* @_T04main13SimpleStorageCMf, i32 0, i32 2) to %swift.type*)

define protected i32 @main(i32, i8**) #0 {
entry:
  %2 = bitcast i8** %1 to i8*
  ret i32 0
}

define hidden swiftcc { i64, i64, i64 } @_T04main13SimpleStorageC10printHelloSSyF(%T4main13SimpleStorageC* swiftself) #0 {
entry:
  %1 = call swiftcc { i64, i64, i64 } @_T0S2SBp21_builtinStringLiteral_Bw17utf8CodeUnitCountBi1_7isASCIItcfC(i8* getelementptr inbounds ([14 x i8], [14 x i8]* @0, i64 0, i64 0), i64 13, i1 true)
  %2 = extractvalue { i64, i64, i64 } %1, 0
  %3 = extractvalue { i64, i64, i64 } %1, 1
  %4 = extractvalue { i64, i64, i64 } %1, 2
  %5 = insertvalue { i64, i64, i64 } undef, i64 %2, 0
  %6 = insertvalue { i64, i64, i64 } %5, i64 %3, 1
  %7 = insertvalue { i64, i64, i64 } %6, i64 %4, 2
  ret { i64, i64, i64 } %7
}

declare swiftcc { i64, i64, i64 } @_T0S2SBp21_builtinStringLiteral_Bw17utf8CodeUnitCountBi1_7isASCIItcfC(i8*, i64, i1) #0

define hidden swiftcc %swift.refcounted* @_T04main13SimpleStorageCfd(%T4main13SimpleStorageC* swiftself) #0 {
entry:
  %1 = bitcast %T4main13SimpleStorageC* %0 to %swift.refcounted*
  ret %swift.refcounted* %1
}

define hidden swiftcc void @_T04main13SimpleStorageCfD(%T4main13SimpleStorageC* swiftself) #0 {
entry:
  %1 = call swiftcc %swift.refcounted* @_T04main13SimpleStorageCfd(%T4main13SimpleStorageC* swiftself %0)
  %2 = bitcast %swift.refcounted* %1 to %T4main13SimpleStorageC*
  %3 = bitcast %T4main13SimpleStorageC* %2 to %swift.refcounted*
  call void @swift_deallocClassInstance(%swift.refcounted* %3, i64 16, i64 7)
  ret void
}

declare void @swift_deallocClassInstance(%swift.refcounted*, i64, i64)

define hidden swiftcc %T4main13SimpleStorageC* @_T04main13SimpleStorageCACycfC(%swift.type* swiftself) #0 {
entry:
  %1 = call %swift.type* @_T04main13SimpleStorageCMa() #3
  %2 = call noalias %swift.refcounted* @swift_rt_swift_allocObject(%swift.type* %1, i64 16, i64 7) #4
  %3 = bitcast %swift.refcounted* %2 to %T4main13SimpleStorageC*
  %4 = call swiftcc %T4main13SimpleStorageC* @_T04main13SimpleStorageCACycfc(%T4main13SimpleStorageC* swiftself %3)
  ret %T4main13SimpleStorageC* %4
}

; Function Attrs: nounwind readnone
define hidden %swift.type* @_T04main13SimpleStorageCMa() #1 {
entry:
  %0 = load %swift.type*, %swift.type** @_T04main13SimpleStorageCML, align 8
  %1 = icmp eq %swift.type* %0, null
  br i1 %1, label %cacheIsNull, label %cont

cacheIsNull:                                      ; preds = %entry
  store atomic %swift.type* bitcast (i64* getelementptr inbounds (<{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>, <{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>* @_T04main13SimpleStorageCMf, i32 0, i32 2) to %swift.type*), %swift.type** @_T04main13SimpleStorageCML release, align 8
  br label %cont

cont:                                             ; preds = %cacheIsNull, %entry
  %2 = phi %swift.type* [ %0, %entry ], [ bitcast (i64* getelementptr inbounds (<{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>, <{ void (%T4main13SimpleStorageC*)*, i8**, i64, %swift.type*, %swift.opaque*, %swift.opaque*, i64, i32, i32, i32, i16, i16, i32, i32, i64, i8*, { i64, i64, i64 } (%T4main13SimpleStorageC*)*, %T4main13SimpleStorageC* (%T4main13SimpleStorageC*)* }>* @_T04main13SimpleStorageCMf, i32 0, i32 2) to %swift.type*), %cacheIsNull ]
  ret %swift.type* %2
}

; Function Attrs: noinline nounwind
define linkonce_odr hidden %swift.refcounted* @swift_rt_swift_allocObject(%swift.type*, i64, i64) #2 {
entry:
  %load = load %swift.refcounted* (%swift.type*, i64, i64)*, %swift.refcounted* (%swift.type*, i64, i64)** @_swift_allocObject
  %3 = tail call %swift.refcounted* %load(%swift.type* %0, i64 %1, i64 %2)
  ret %swift.refcounted* %3
}

define hidden swiftcc %T4main13SimpleStorageC* @_T04main13SimpleStorageCACycfc(%T4main13SimpleStorageC* swiftself) #0 {
entry:
  ret %T4main13SimpleStorageC* %0
}

define private %swift.type** @get_field_types_SimpleStorage(%swift.type* %SimpleStorage) #0 {
entry:
  %0 = load %swift.type**, %swift.type*** @field_type_vector_SimpleStorage, align 8
  %1 = icmp eq %swift.type** %0, null
  br i1 %1, label %build_field_types, label %done

build_field_types:                                ; preds = %entry
  %2 = call noalias i8* @swift_rt_swift_slowAlloc(i64 0, i64 7) #4
  %3 = bitcast i8* %2 to %swift.type**
  %4 = ptrtoint %swift.type** %3 to i64
  %5 = cmpxchg i64* bitcast (%swift.type*** @field_type_vector_SimpleStorage to i64*), i64 0, i64 %4 seq_cst seq_cst
  %6 = extractvalue { i64, i1 } %5, 1
  %7 = extractvalue { i64, i1 } %5, 0
  br i1 %6, label %done, label %race_lost

race_lost:                                        ; preds = %build_field_types
  call void @swift_rt_swift_slowDealloc(i8* %2, i64 0, i64 7) #4
  %8 = inttoptr i64 %7 to %swift.type**
  br label %done

done:                                             ; preds = %race_lost, %build_field_types, %entry
  %9 = phi %swift.type** [ %0, %entry ], [ %3, %build_field_types ], [ %8, %race_lost ]
  ret %swift.type** %9
}

; Function Attrs: noinline nounwind
define linkonce_odr hidden i8* @swift_rt_swift_slowAlloc(i64, i64) #2 {
entry:
  %load = load i8* (i64, i64)*, i8* (i64, i64)** @_swift_slowAlloc
  %2 = tail call i8* %load(i64 %0, i64 %1)
  ret i8* %2
}

; Function Attrs: noinline nounwind
define linkonce_odr hidden void @swift_rt_swift_slowDealloc(i8*, i64, i64) #2 {
entry:
  %load = load void (i8*, i64, i64)*, void (i8*, i64, i64)** @_swift_slowDealloc
  tail call void %load(i8* %0, i64 %1, i64 %2)
  ret void
}

attributes #0 = { "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "target-cpu"="x86-64" "target-features"="+fxsr,+mmx,+sse,+sse2,+x87" }
attributes #1 = { nounwind readnone "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "target-cpu"="x86-64" "target-features"="+fxsr,+mmx,+sse,+sse2,+x87" }
attributes #2 = { noinline nounwind }
attributes #3 = { nounwind readnone }
attributes #4 = { nounwind }

!llvm.module.flags = !{!0, !2, !3}

!0 = !{i32 6, !"Linker Options", !1}
!1 = !{}
!2 = !{i32 4, !"Objective-C Garbage Collection", i32 1280}
!3 = !{i32 1, !"Swift Version", i32 5}
