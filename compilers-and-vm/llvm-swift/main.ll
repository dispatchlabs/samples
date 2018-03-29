; ModuleID = '-'
source_filename = "-"
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

%struct._SwiftEmptyArrayStorage = type { %struct.HeapObject, %struct._SwiftArrayBodyStorage }
%struct.HeapObject = type { %struct.HeapMetadata*, %struct.InlineRefCounts }
%struct.HeapMetadata = type opaque
%struct.InlineRefCounts = type { i32, i32 }
%struct._SwiftArrayBodyStorage = type { i64, i64 }
%Ts6UInt32V = type <{ i32 }>
%swift.refcounted = type { %swift.type*, i32, i32 }
%swift.type = type { i64 }
%Ts27_ContiguousArrayStorageBaseC = type opaque
%Any = type { [24 x i8], %swift.type* }
%TSS = type <{ %Ts11_StringCoreV }>
%Ts11_StringCoreV = type <{ %TSvSg, %TSu, %TyXlSg }>
%TSvSg = type <{ [8 x i8] }>
%TSu = type <{ i64 }>
%TyXlSg = type <{ [8 x i8] }>

@_swiftEmptyArrayStorage = external global %struct._SwiftEmptyArrayStorage, align 8
@_T0s19_emptyStringStorages6UInt32Vv = external global %Ts6UInt32V, align 4
@_swift_retain = external global void (%swift.refcounted*)*
@_swift_release = external global void (%swift.refcounted*)*
@_T0SSN = external global %swift.type, align 8
@0 = private unnamed_addr constant [14 x i8] c"Hello, World!\00"
@__swift_reflection_version = linkonce_odr hidden constant i16 3
@_swift1_autolink_entries = private constant [37 x i8] c"-lswiftCore\00-lswiftSwiftOnoneSupport\00", section ".swift1_autolink_entries", align 8
@llvm.used = appending global [2 x i8*] [i8* bitcast (i16* @__swift_reflection_version to i8*), i8* getelementptr inbounds ([37 x i8], [37 x i8]* @_swift1_autolink_entries, i32 0, i32 0)], section "llvm.metadata", align 8

define protected i32 @main(i32, i8**) #0 {
entry:
  %2 = bitcast i8** %1 to i8*
  call swiftcc void @_T04mainAAyyF()
  ret i32 0
}

define hidden swiftcc void @_T04mainAAyyF() #0 {
entry:
  %0 = call swiftcc { %Ts27_ContiguousArrayStorageBaseC*, i8* } @_T0s27_allocateUninitializedArraySayxG_BptBwlFyp_Tgq5(i64 1)
  %1 = extractvalue { %Ts27_ContiguousArrayStorageBaseC*, i8* } %0, 0
  %2 = extractvalue { %Ts27_ContiguousArrayStorageBaseC*, i8* } %0, 1
  %3 = bitcast %Ts27_ContiguousArrayStorageBaseC* %1 to %swift.refcounted*
  call void @swift_rt_swift_retain(%swift.refcounted* %3) #3
  call void bitcast (void (%swift.refcounted*)* @swift_rt_swift_release to void (%Ts27_ContiguousArrayStorageBaseC*)*)(%Ts27_ContiguousArrayStorageBaseC* %1) #3
  %4 = bitcast i8* %2 to %Any*
  %5 = getelementptr inbounds %Any, %Any* %4, i32 0, i32 1
  store %swift.type* @_T0SSN, %swift.type** %5, align 8
  %6 = getelementptr inbounds %Any, %Any* %4, i32 0, i32 0
  %7 = getelementptr inbounds %Any, %Any* %4, i32 0, i32 0
  %8 = bitcast [24 x i8]* %7 to %TSS*
  %9 = call swiftcc { i64, i64, i64 } @_T0S2SBp21_builtinStringLiteral_Bw17utf8CodeUnitCountBi1_7isASCIItcfC(i8* getelementptr inbounds ([14 x i8], [14 x i8]* @0, i64 0, i64 0), i64 13, i1 true)
  %10 = extractvalue { i64, i64, i64 } %9, 0
  %11 = extractvalue { i64, i64, i64 } %9, 1
  %12 = extractvalue { i64, i64, i64 } %9, 2
  %._core = getelementptr inbounds %TSS, %TSS* %8, i32 0, i32 0
  %._core._baseAddress = getelementptr inbounds %Ts11_StringCoreV, %Ts11_StringCoreV* %._core, i32 0, i32 0
  %13 = bitcast %TSvSg* %._core._baseAddress to i64*
  store i64 %10, i64* %13, align 8
  %._core._countAndFlags = getelementptr inbounds %Ts11_StringCoreV, %Ts11_StringCoreV* %._core, i32 0, i32 1
  %._core._countAndFlags._value = getelementptr inbounds %TSu, %TSu* %._core._countAndFlags, i32 0, i32 0
  store i64 %11, i64* %._core._countAndFlags._value, align 8
  %._core._owner = getelementptr inbounds %Ts11_StringCoreV, %Ts11_StringCoreV* %._core, i32 0, i32 2
  %14 = bitcast %TyXlSg* %._core._owner to i64*
  store i64 %12, i64* %14, align 8
  %15 = call swiftcc { i64, i64, i64 } @_T0s5printySayypGd_SS9separatorSS10terminatortFfA0_()
  %16 = extractvalue { i64, i64, i64 } %15, 0
  %17 = extractvalue { i64, i64, i64 } %15, 1
  %18 = extractvalue { i64, i64, i64 } %15, 2
  %19 = call swiftcc { i64, i64, i64 } @_T0s5printySayypGd_SS9separatorSS10terminatortFfA1_()
  %20 = extractvalue { i64, i64, i64 } %19, 0
  %21 = extractvalue { i64, i64, i64 } %19, 1
  %22 = extractvalue { i64, i64, i64 } %19, 2
  call swiftcc void @_T0s5printySayypGd_SS9separatorSS10terminatortF(%Ts27_ContiguousArrayStorageBaseC* %1, i64 %16, i64 %17, i64 %18, i64 %20, i64 %21, i64 %22)
  ret void
}

; Function Attrs: noinline
declare swiftcc void @_T0s5printySayypGd_SS9separatorSS10terminatortF(%Ts27_ContiguousArrayStorageBaseC*, i64, i64, i64, i64, i64, i64) #1

declare swiftcc { %Ts27_ContiguousArrayStorageBaseC*, i8* } @_T0s27_allocateUninitializedArraySayxG_BptBwlFyp_Tgq5(i64) #0

; Function Attrs: noinline nounwind
define linkonce_odr hidden void @swift_rt_swift_retain(%swift.refcounted*) #2 {
entry:
  %load = load void (%swift.refcounted*)*, void (%swift.refcounted*)** @_swift_retain
  tail call void %load(%swift.refcounted* %0)
  ret void
}

; Function Attrs: noinline nounwind
define linkonce_odr hidden void @swift_rt_swift_release(%swift.refcounted*) #2 {
entry:
  %load = load void (%swift.refcounted*)*, void (%swift.refcounted*)** @_swift_release
  tail call void %load(%swift.refcounted* %0)
  ret void
}

declare swiftcc { i64, i64, i64 } @_T0S2SBp21_builtinStringLiteral_Bw17utf8CodeUnitCountBi1_7isASCIItcfC(i8*, i64, i1) #0

; Function Attrs: noinline
declare swiftcc { i64, i64, i64 } @_T0s5printySayypGd_SS9separatorSS10terminatortFfA0_() #1

; Function Attrs: noinline
declare swiftcc { i64, i64, i64 } @_T0s5printySayypGd_SS9separatorSS10terminatortFfA1_() #1

attributes #0 = { "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "target-cpu"="x86-64" "target-features"="+fxsr,+mmx,+sse,+sse2,+x87" }
attributes #1 = { noinline "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "target-cpu"="x86-64" "target-features"="+fxsr,+mmx,+sse,+sse2,+x87" }
attributes #2 = { noinline nounwind }
attributes #3 = { nounwind }

!llvm.module.flags = !{!0, !2, !3}

!0 = !{i32 6, !"Linker Options", !1}
!1 = !{}
!2 = !{i32 4, !"Objective-C Garbage Collection", i32 1280}
!3 = !{i32 1, !"Swift Version", i32 5}
